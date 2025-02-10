# Verificação de Qualidade - Rappi Turbo 🚀

Este repositório contém a estrutura de documentação e testes automatizados para verificar a precisão da tela de ganhos e a eficiência do onboarding no aplicativo Rappi Turbo.

## 📌 Business Drivers (Mapa)
### **DN1 - Precisão da Tela de Ganhos**
- **Requisito**: A tela de ganhos do entregador deve exibir informações corretas e atualizadas em 99,9% das consultas.
- **Métrica**: O atraso na exibição dos valores não pode ultrapassar 5 segundos em 95% das requisições.
- **Monitoramento**: Se houver discrepâncias entre os valores mostrados e os valores reais processados no backend, um alerta deve ser disparado.

### **DN2 - Eficiência do Onboarding**
- **Requisito**: O fluxo de onboarding deve ser concluído em menos de 5 minutos para 90% dos novos entregadores.
- **Métrica**: A taxa de conclusão do onboarding deve ser maior que 85%. Caso contrário, identificar pontos de abandono.
- **Monitoramento**: O sistema deve detectar e registrar os momentos de maior abandono dentro do processo.

## 📜 Cenários de Teste (Gherkin)

### **Testes para DN1 - Precisão da Tela de Ganhos**
```gherkin
#language: pt
Funcionalidade: Precisão da Tela de Ganhos

  Cenário: Exibição correta dos ganhos do entregador
    Dado que um entregador tem ganhos processados no backend
    Quando ele acessa a tela de ganhos
    Então os valores exibidos devem corresponder aos valores reais armazenados

  Cenário: Atraso na atualização dos ganhos
    Dado que o backend processou uma atualização nos ganhos do entregador
    Quando a tela de ganhos for carregada
    Então a atualização deve ocorrer em menos de 5 segundos

  Cenário: Disparo de alerta em caso de discrepância
    Dado que há uma diferença entre os ganhos exibidos e os ganhos reais
    Então um alerta deve ser enviado ao sistema de monitoramento
```

### **Testes para DN2 - Eficiência do Onboarding**
```gherkin
#language: pt
Funcionalidade: Eficiência do Onboarding

  Cenário: Novo entregador completa o onboarding em menos de 5 minutos
    Dado que um novo entregador inicia o processo de onboarding
    Quando ele conclui todas as etapas
    Então o tempo total não deve exceder 5 minutos

  Cenário: Identificação de abandono durante o onboarding
    Dado que um entregador iniciou o onboarding
    Quando ele abandona o processo antes de completar
    Então o sistema deve registrar o momento do abandono

  Cenário: Taxa de conclusão mínima de 85%
    Dado que novos entregadores estão realizando o onboarding
    Quando a taxa de conclusão for calculada
    Então ela deve ser de pelo menos 85%
    E se for menor, os pontos de abandono devem ser analisados
```

## 🔍 Testes Automatizados em Go

### **Testes para DN1 - Precisão da Tela de Ganhos**
```go
package main

import (
	"testing"
	"time"
)

// Mock de ganhos no backend
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
```

### **Testes para DN2 - Eficiência do Onboarding**
```go
package main

import (
	"testing"
	"time"
)

// Simula o tempo de onboarding de um entregador
func tempoOnboarding() time.Duration {
	return 4 * time.Minute // Simulação de um onboarding dentro do limite
}

// Teste: Verificar se o onboarding ocorre em menos de 5 minutos
func TestTempoOnboarding(t *testing.T) {
	tempo := tempoOnboarding()
	if tempo > 5*time.Minute {
		t.Errorf("Erro: tempo de onboarding (%.2f minutos) acima do limite de 5 minutos", tempo.Minutes())
	}
}

// Simula taxa de conclusão do onboarding
func taxaConclusaoOnboarding() float64 {
	return 88.0 // Simulação de taxa acima de 85%
}

// Teste: Verificar se a taxa de conclusão é de pelo menos 85%
func TestTaxaConclusaoOnboarding(t *testing.T) {
	taxa := taxaConclusaoOnboarding()
	if taxa < 85.0 {
		t.Errorf("Erro: taxa de conclusão do onboarding é %.2f%%, abaixo do limite de 85%%", taxa)
	}
}
```
