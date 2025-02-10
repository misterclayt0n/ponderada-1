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
Para rodar, digite no terminal:

```bash
go test ./...
```
