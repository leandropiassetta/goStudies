package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evaluateProximityRule(t *testing.T) {

	sample := "joazinho tem o cpf de numero 16538533 aonde tem cpf tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 10,
		windowAfter:  15,
		global:       0,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.True(t, result)

}

func Test_evaluateProximityRule_Validations_Window(t *testing.T) {

	sample := "joazinho tem o cpf de numero 16538533 aonde tem cpf tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 60,
		windowAfter:  60,
		global:       0,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.True(t, result)

}

func Test_evaluateProximityRule_Not_Windows(t *testing.T) {

	sample := "joazinho tem o cpf de numero 16538533 aonde tem cpf tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 0,
		windowAfter:  0,
		global:       0,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.False(t, result)
}

func Test_evaluateProximityRule_False(t *testing.T) {

	sample := "joazinho tem o cpf de numero 16538533 aonde tem cpf tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       0,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.False(t, result)

}

func Test_evaluateProximityRule_Global(t *testing.T) {

	sample := "joazinho tem o cpf de numero 16538533 aonde tem cpf tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       1,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.True(t, result)

}

func Test_evaluateProximityRule_Global_False(t *testing.T) {

	sample := "joazinho tem o de numero 16538533 aonde tem tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       1,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.False(t, result)

}

func Test_evaluateProximityRule_Global_(t *testing.T) {

	sample := "joazinho tem o de numero 16538533 aonde tem tambem"
	pattern := regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)")
	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       1,
	}

	findingLocation := []int{29, 36}

	result := evaluateProximityRule(sample, *pattern, findingLocation, window)

	assert.False(t, result)

}

func Test_evaluateFindingByProximityRules_Relative(t *testing.T) {
	sample := "joazinho tem o de cpf numero 16538533 aonde tem tambem"

	findings := []Finding{
		{
			Quote:      "16538533",
			Location:   []int{25, 33},
			Likelihood: 1,
		},
	}

	likelihoodAdjustment := LikelihoodAdjustment{
		relativeLikelihood: 2,
		fixedLikelihood:    0,
	}

	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       0,
	}

	rules := []Rule{
		{
			pattern:              regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)"),
			window:               window,
			likelihoodAdjustment: likelihoodAdjustment,
		},
	}

	channel := make(chan Finding, len(findings))

	evaluateFindingByProximityRules(findings[0], sample, rules, channel)

	result := <-channel

	assert.Equal(t, Finding{Quote: "16538533", Location: []int{25, 33}, Likelihood: 3}, result)
}

func Test_evaluateFindingByProximityRules_Fixed(t *testing.T) {
	sample := "joazinho tem o de cpf numero 16538533 aonde tem tambem"

	findings := []Finding{
		{
			Quote:      "16538533",
			Location:   []int{25, 33},
			Likelihood: 1,
		},
	}

	likelihoodAdjustment := LikelihoodAdjustment{
		relativeLikelihood: 0,
		fixedLikelihood:    2,
	}

	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       0,
	}

	rules := []Rule{
		{
			pattern:              regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)"),
			window:               window,
			likelihoodAdjustment: likelihoodAdjustment,
		},
	}

	channel := make(chan Finding, len(findings))

	evaluateFindingByProximityRules(findings[0], sample, rules, channel)

	result := <-channel

	assert.Equal(t, Finding{Quote: "16538533", Location: []int{25, 33}, Likelihood: 2}, result)
}

func Test_evaluateFindingsByProximityRules(t *testing.T) {
	sample := "joazinho tem o de cpf numero 16538533 aonde tem tambem"

	findings := []Finding{
		{
			Quote:      "16538533",
			Location:   []int{25, 33},
			Likelihood: 1,
		},
	}

	likelihoodAdjustment := LikelihoodAdjustment{
		relativeLikelihood: 0,
		fixedLikelihood:    2,
	}

	window := Window{
		windowBefore: 10,
		windowAfter:  5,
		global:       0,
	}

	rules := []Rule{
		{
			pattern:              regexp.MustCompile("(cpf|CPF|Cpf|c.p.f)"),
			window:               window,
			likelihoodAdjustment: likelihoodAdjustment,
		},
	}

	channel := make(chan Finding, len(findings))

	evaluateFindingsByProximityRules(findings, sample, rules, channel)

	result := <-channel

	assert.Equal(t, Finding{Quote: "16538533", Location: []int{25, 33}, Likelihood: 2}, result)
}
