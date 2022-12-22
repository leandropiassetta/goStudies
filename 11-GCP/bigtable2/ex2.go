package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
)

const (
	tableName        = ""
	columnFamilyName = ""
	columnName       = ""
)

func main() {
	project := ""
	instance := ""

	ctx := context.Background()

	creds, _ := base64.StdEncoding.DecodeString("")

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

	tblInfo, err := adminClient.TableInfo(ctx, tableName)

	fmt.Println("Tabela Info ==>", tblInfo.FamilyInfos)

	if err != nil {
		log.Fatalf("Could not read info for table %s: %v", tableName, err)
	}

	// Precisa de credenciais tanto pra ler como pra escrever.

	writeSimple(project, instance, tableName, creds)

	err = readRows(project, instance, tableName, creds)

	fmt.Print(err)

	//

	/*if err := adminClient.CreateColumnFamily(ctx, tableName, "Datos_Personales"); err != nil {
		log.Fatalf("Could not create column family %s: %v", columnFamilyName, err)
	}

	*/
}

func writeSimple(projectID, instanceID, tableName string, creds []byte) error {
	ctx := context.Background()

	client, err := bigtable.NewClient(ctx, projectID, instanceID, option.WithCredentialsJSON(creds))
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}

	defer client.Close()

	tbl := client.Open(tableName)
	columnFamilyName := "Datos_Personales"
	timestamp := bigtable.Now()

	// vai modificar a tabela
	mut := bigtable.NewMutation()

	// Seta colunas pai e filhas, data e valor
	mut.Set(columnFamilyName, "Name", timestamp, []byte("Romario"))
	mut.Set(columnFamilyName, "Age", timestamp, []byte("60"))
	mut.Set(columnFamilyName, "email", timestamp, []byte("paulobaier@gmail.com"))
	mut.Set(columnFamilyName, "cpf", timestamp, []byte("00011122233"))
	mut.Set(columnFamilyName, "telefone", timestamp, []byte("(048)32866456"))

	rowKey := "1004"

	if err := tbl.Apply(ctx, rowKey, mut); err != nil {
		return fmt.Errorf("apply: %v", err)
	}
	fmt.Printf("Successfully wrote row: %s\n", rowKey)

	return nil
}

// x, _ := json.Marshal(rowData)

// fmt.Println("JSON", string(x))

// fmt.Println("ROWDATA ===>", string(x))

func readRows(projectID, instanceID string, tableName string, creds []byte) error {
	ctx := context.Background()

	client, err := bigtable.NewClient(ctx, projectID, instanceID, option.WithCredentialsJSON(creds))
	if err != nil {
		return fmt.Errorf("bigtable.NewClient: %v", err)
	}

	defer client.Close()

	// filter := bigtable.ChainFilters(bigtable.LatestNFilter(1), bigtable.FamilyFilter("Datos_Personales"), bigtable.ColumnFilter("Name"))

	// abre a tabela, ler ou escrever
	tbl := client.Open(tableName)

	// rowKey := "age#60"
	// row, _ := tbl.ReadRow(ctx, rowKey)

	// fmt.Println(row)

	err = tbl.ReadRows(ctx, bigtable.RowRange{},
		func(row bigtable.Row) bool {
			printRow(row)
			return true
			// RowFilter -> retorna opcoes para aplicar em filter, combinar filtros use ChainFilter
		})
	// bigtable.RowFilter(filter)
	if err != nil {
		return fmt.Errorf("error table.NewClient: %v", err)
	}

	return nil
}

func printRow(row bigtable.Row) {
	fmt.Printf("Reading data for %s:\n", row.Key())

	for columnFamily, cols := range row {
		fmt.Printf("Column Family %s\n", columnFamily)

		for _, col := range cols {
			// qualifier -> name
			qualifier := col.Column[strings.IndexByte(col.Column, ':')+1:]

			fmt.Printf("\t%s: %s @%d\n", qualifier, col.Value, col.Timestamp)
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
	filter := bigtable.ColumnRangeFilter("Datos_Personales", "Age", "")

	return readWithFilter(projectID, instanceID, tableName, filter)
}
