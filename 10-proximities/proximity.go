package main

import (
	"fmt"
	"regexp"
)

type Finding struct {
	Quote      string
	Location   []int
	Likelihood int
}

type Proximity struct {
	datatype string
	rules    []Rule
}

type Rule struct {
	pattern              *regexp.Regexp
	window               Window
	likelihoodAdjustment LikelihoodAdjustment
}

type Window struct {
	windowBefore int
	windowAfter  int
	global       int
}

type LikelihoodAdjustment struct {
	relativeLikelihood int
	fixedLikelihood    int
}

func evaluateProximityRule(text string, pattern regexp.Regexp, location []int, window Window) bool {
	if window.global == 1 {
		return pattern.MatchString(text)
	}

	if window.windowBefore == 0 && window.windowAfter == 0 {
		return false
	}

	sliceBegin := location[0] - window.windowBefore

	if sliceBegin < 0 {
		sliceBegin = 0
	}

	sliceAfter := location[1] + window.windowAfter

	if sliceAfter > len(text) {
		sliceAfter = len(text)
	}

	sliceText := text[sliceBegin:location[0]] + text[location[1]:sliceAfter]

	hasProximity := pattern.MatchString(sliceText)

	return hasProximity
}

func evaluateFindingByProximityRules(finding Finding, sample string, rules []Rule, channelResult chan<- Finding) {
	newLikelihood := finding.Likelihood

	for _, rule := range rules {
		window := rule.window
		pattern := rule.pattern

		relative := rule.likelihoodAdjustment.relativeLikelihood
		fixed := rule.likelihoodAdjustment.fixedLikelihood

		if evaluateProximityRule(sample, *pattern, finding.Location, window) {
			if relative != 0 {
				newLikelihood += relative
			}

			if fixed != 0 {
				newLikelihood = fixed
			}
		}
	}

	finding.Likelihood = newLikelihood

	channelResult <- finding
}

// finishedTask := 0

// for result := range channelResult {
// 	finishedTask++
// 	fmt.Println(result)
// 	if finishedTask == len(findings) {
// 		close(channelResult)
// 	}
// }

func evaluateFindingsByProximityRules(findings []Finding, sample string, rules []Rule, channelResult chan<- Finding) {
	for _, finding := range findings {
		go evaluateFindingByProximityRules(finding, sample, rules, channelResult)
	}
}

func main() {
	findings := []Finding{
		{
			Quote:      "8989999919",
			Location:   []int{14, 22},
			Likelihood: 0,
		},
		{
			Quote:      "34567898991",
			Location:   []int{70, 81},
			Likelihood: 1,
		},
		{
			Quote:      "34567892139",
			Location:   []int{150, 161},
			Likelihood: 1,
		},
	}

	exWindow := Window{
		windowBefore: 15,
		windowAfter:  0,
		global:       0,
	}

	exWindow2 := Window{
		windowBefore: 15,
		windowAfter:  25,
		global:       0,
	}

	// exWindow3 := Window{
	// 	windowBefore: 15,
	// 	windowAfter:  0,
	// 	global:       0,
	// }

	exlikelihoodAdjustment := LikelihoodAdjustment{
		relativeLikelihood: 0,
		fixedLikelihood:    2,
	}

	exlikelihoodAdjustment2 := LikelihoodAdjustment{
		relativeLikelihood: 1,
		fixedLikelihood:    0,
	}

	exRules := Rule{
		pattern:              regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)"),
		window:               exWindow,
		likelihoodAdjustment: exlikelihoodAdjustment,
	}

	exRules2 := Rule{
		pattern:              regexp.MustCompile("(federal|FEDERAL)"),
		window:               exWindow2,
		likelihoodAdjustment: exlikelihoodAdjustment2,
	}

	// exRules3 := Rule{
	// 	pattern:              regexp.MustCompile("(federal|FEDERAL)"),
	// 	window:               exWindow3,
	// 	likelihoodAdjustment: exlikelihoodAdjustment2,
	// }

	proximity1 := []Proximity{
		{
			datatype: "CPF",
			rules:    []Rule{exRules},
		},
		{
			datatype: "Email",
			rules:    []Rule{exRules2},
		},
	}

	fmt.Println("STRUCT PROXIMITY: ->", proximity1)
	fmt.Println("------------------------------------")

	// COORDENADAS
	// 14,24 // 70,81 //152, 163
	sample := "João tem o cpf 8989999919 no GOVERNO federal, fernando tem outro CPF nº 34567898991 na policia FEDERAL, Fernandinha passou na federal e seu cpf é 34567892139."

	channelResult := make(chan Finding, len(findings))

	evaluateFindingsByProximityRules(findings, sample, proximity1[0].rules, channelResult)

	finishedTask := 0

	for result := range channelResult {
		finishedTask++
		fmt.Println(result)
		if finishedTask == len(findings) {
			close(channelResult)
		}
	}
	fmt.Println("------------------------------------")
}

// 69 - 70

// match_use_case

/*
	//DATATYPE VÁLIDO

	//validar se es un datatype válido, volver un error en caso de que sean repetidos los datatypes
	// va volver las proximities, devolver un mapa que la clave es dataype y value las proximities, retornar un error...
*/

// 123 - scanner

// 	// --> validar regex de las proximities, y ya compila *regex

// scannerValidateProximityRules(request detail rules) {
// 	// esta funcion tiene que validar las entradas proximities de request dto, y converted a un domain de rules, con el que yo necesite, y el int de los likelihood

// }
