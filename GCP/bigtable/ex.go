package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"sort"
	"strings"

	"cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
)

const (
	tableName        = "Criciuma"
	columnFamilyName = "Jogadores"
	columnName       = "Posicao"
)

var jogadores = []string{"Alexandre", "Paulo Baier", "Wilsão", "Silvio Criciúma", "Clebér Gaúcho", "Itá", "Mahicon Librelato", "Roberto Cavalo", "Soares", "Grizzo", "Jairo Lenzi"}

func sliceContains(list []string, target string) bool {
	for _, s := range list {
		if s == target {
			return true
		}
	}
	return false
}

func printRow(row bigtable.Row) {
	fmt.Println("Reading data for %s:\n", row.Key())
	var keys []string
	for k := range row {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, columnFamily := range keys {
		cols := row[columnFamily]
		fmt.Println("Column Family %s\n", columnFamily)

		for _, col := range cols {
			qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]
			fmt.Println("\t%s: %s @%d\n", qualifier, col.Value, col.Timestamp)
		}
	}
}

func readWithFilter(projectID, instanceID string, tableName string, filter bigtable.Filter) error {

	ctx := context.Background()
	client, err := bigtable.NewClient(ctx, projectID, instanceID)
	if err != nil {
		return fmt.Errorf("bigtable.NewAdminClient: %v", err)
	}
	tbl := client.Open(tableName)
	_ = tbl.ReadRows(ctx, bigtable.RowRange{},
		func(row bigtable.Row) bool {
			printRow(row)
			return true
		}, bigtable.RowFilter(filter))

	if err = client.Close(); err != nil {
		return fmt.Errorf("client.Close(): %v", err)
	}

	return nil
}

func filterLimitColRange(projectID, instanceID string, tableName string) error {
	filter := bigtable.ColumnRangeFilter("Jogadores", "0", "10")

	return readWithFilter(projectID, instanceID, tableName, filter)
}

func main() {
	project := "datasec-discovery"
	instance := "discovery-test-2"
	projectID := "datasec-discovery"
	instanceID := "discovery-test-2"

	ctx := context.Background()

	//credenciais
	creds, _ := base64.StdEncoding.DecodeString()

	// Para gerenciar tabelas, conecte-se ao Bigtable usando bigtable.NewAdminClient().

	adminClient, err := bigtable.NewAdminClient(ctx, project, instance, option.WithCredentialsJSON(creds))
	if err != nil {
		log.Fatalf("Could not create admin client: %v", err)
	}

	// Retorna a lista de tabelas na instância

	tables, err := adminClient.Tables(ctx)
	if err != nil {
		log.Fatalf("Could not fetch table list: %v", err)
	}

	fmt.Println("TABLES", tables)

	// Crie uma tabela com AdminClient.CreateTable() e receba informações sobre ela com AdminClient.TableInfo(). Crie um grupo de colunas com AdminClient.CreateColumnFamily().
	if !sliceContains(tables, tableName) {
		log.Printf("Creating table %s", tableName)
		if err := adminClient.CreateTable(ctx, tableName); err != nil {
			log.Fatalf("Could not create table %s: %v", tableName, err)
		}
	}

	tblInfo, err := adminClient.TableInfo(ctx, tableName)

	for _, columns := range tblInfo.FamilyInfos {
		fmt.Println("COLUMNS: ", columns.Name)
	}

	fmt.Println("Tabela Info ==>", tblInfo.FamilyInfos)
	if err != nil {
		log.Fatalf("Could not read info for table %s: %v", tableName, err)
	}

	if !sliceContains(tblInfo.Families, columnFamilyName) {
		if err := adminClient.CreateColumnFamily(ctx, tableName, columnFamilyName); err != nil {
			log.Fatalf("Could not create column family %s: %v", columnFamilyName, err)
		}
	}

	// Como gravar linhas em uma tabela

	// Abra a tabela em que você quer gravar. Use bigtable.NewMutation() para criar uma mutação em uma única linha e use Mutation.Set() para definir valores na linha. Gere uma chave de linha exclusiva para cada linha. Repita esses passos para criar várias mutações. Por fim, use Table.ApplyBulk() para aplicar todas as mutações à tabela.

	client, _ := bigtable.NewClient(ctx, project, instance, option.WithCredentialsJSON(creds))
	table := client.Open(tableName)
	muts := make([]*bigtable.Mutation, len(jogadores))
	rowKeys := make([]string, len(jogadores))

	log.Printf("Writing jogadores rows to table")
	for i, jogador := range jogadores {
		muts[i] = bigtable.NewMutation()
		muts[i].Set(columnFamilyName, columnName, bigtable.Now(), []byte(jogador))

		rowKeys[i] = fmt.Sprintf("%s %d", columnName, i+1)
	}

	rowErrs, _ := table.ApplyBulk(ctx, rowKeys, muts)

	if err != nil {
		log.Fatalf("Could not apply bulk row mutation: %v", err)
	}
	if rowErrs != nil {
		for _, rowErr := range rowErrs {
			log.Printf("Error writing row: %v", rowErr)
		}
		log.Fatalf("Could not write some rows")
	}

	log.Printf("Getting a single jogador by row key:")
	row, err := table.ReadRow(ctx, rowKeys[0], bigtable.RowFilter(bigtable.ColumnFilter(columnName)))
	if err != nil {
		log.Fatalf("Could not read row with key %s: %v", rowKeys[0], err)
	}
	log.Printf("\t%s = %s\n", rowKeys[0], string(row[columnFamilyName][0].Value))

	// Use Table.ReadRows() para verificar todas as linhas em uma tabela. Feche o cliente de dados quando terminar de usá-lo.

	log.Printf("Reading all jogadores rows:")
	_ = table.ReadRows(ctx, bigtable.PrefixRange(columnName), func(row bigtable.Row) bool {
		item := row[columnFamilyName][0]

		log.Printf("\t%s = %s\n", item.Row, string(item.Value))
		return true
	}, bigtable.RowFilter(bigtable.ColumnFilter(columnName)))

	if err = adminClient.Close(); err != nil {
		log.Fatalf("Could not close data operations client: %v", err)
	}

	filterLimitColRange(projectID, instanceID, tableName)

	// //Exclua uma tabela com AdminClient.DeleteTable(). Feche o cliente de administração quando terminar de usá-lo.

	// log.Printf("Deleting the table")
	// if err = adminClient.DeleteTable(ctx, tableName); err != nil {
	// 	log.Fatalf("Could not delete table %s: %v", tableName, err)
	// }

	// if err = adminClient.Close(); err != nil {
	// 	log.Fatalf("Could not close admin client: %v", err)
	// }

}

//Objetivos.

// Listar las tablas de una instancia

// Listar las columnas de una tabla

// Listar los datos de una columna

// Listas las instancias de un proyecto

// Como listar las instancias de big table que tiene un proyecto.

// Verificar los permisos de services account.
