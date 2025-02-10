package main

import (
	"testing"
	"time"
)

var backendGanhos = 100.00

func exibirGanhos() float64 {
	time.Sleep(1 * time.Second)
	return backendGanhos
}

func TestPrecisaoTelaGanhos(t *testing.T) {
	ganhosApp := exibirGanhos()
	if ganhosApp != backendGanhos {
		t.Errorf("Erro: ganhos exibidos (%.2f) não correspondem ao backend (%.2f)", ganhosApp, backendGanhos)
	}
}

func TestTempoAtualizacaoGanhos(t *testing.T) {
	start := time.Now()
	exibirGanhos()
	duration := time.Since(start)

	if duration.Seconds() > 5 {
		t.Errorf("Erro: atualização dos ganhos demorou %.2f segundos, acima do limite de 5s", duration.Seconds())
	}
}

func tempoOnboarding() time.Duration {
	return 4 * time.Minute
}

func TestTempoOnboarding(t *testing.T) {
	tempo := tempoOnboarding()
	if tempo > 5*time.Minute {
		t.Errorf("Erro: tempo de onboarding (%.2f minutos) acima do limite de 5 minutos", tempo.Minutes())
	}
}

func taxaConclusaoOnboarding() float64 {
	return 88.0
}

func TestTaxaConclusaoOnboarding(t *testing.T) {
	taxa := taxaConclusaoOnboarding()
	if taxa < 85.0 {
		t.Errorf("Erro: taxa de conclusão do onboarding é %.2f%%, abaixo do limite de 85%%", taxa)
	}
}
