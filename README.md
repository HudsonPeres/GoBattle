# Go Battle!

## Visão Geral do Projeto

O **Go Battle!** é um jogo de combate por turnos simples, desenvolvido inteiramente em Go. O projeto demonstra conceitos fundamentais da linguagem, como interfaces, structs, métodos e lógica de jogo baseada em consola. Os jogadores podem escolher entre diferentes classes de heróis, cada uma com atributos e ataques únicos, para enfrentar monstros gerados aleatoriamente num combate até à morte.

## Funcionalidades

- **Criação de Herói**: O jogador pode definir o nome do seu personagem.
- **Escolha de Classe**: Duas classes distintas com diferentes níveis de vida e poder de ataque:
  - **Guerreiro**: Possui mais vida (20) e causa dano consistente com a sua espada.
  - **Mago**: Possui menos vida (10), mas tem o potencial de causar danos maiores com as suas bolas de fogo.
- **Combate por Turnos**: Sistema de batalha automatizado onde o jogador e o monstro atacam alternadamente.
- **Geração Aleatória**: A vida do monstro e o dano de cada ataque são gerados aleatoriamente, tornando cada partida única.
- **Interface de Consola**: Interação simples e direta através do terminal com emojis para uma experiência mais visual.

## Estrutura do Código

O código está organizado de forma modular utilizando os seguintes conceitos:

- **Interface `Personagem`**: Define o contrato para qualquer herói no jogo, exigindo métodos para obter nome, vida, atacar e receber dano.
- **Structs `Guerreiro` e `Mago`**: Implementam a interface `Personagem` com comportamentos específicos para cada classe.
- **Struct `Monstro`**: Representa o inimigo com a sua própria lógica de ataque.
- **Função `IniciarCombate`**: Gere o loop principal da batalha, processando os turnos até que um dos combatentes fique sem vida.

## Como Executar o Projeto

Para jogar o **Go Battle!**, siga os passos abaixo:

1.  **Ter o Go instalado** no seu sistema.

2.  **Guarde o código** num ficheiro chamado `Gobattle.go`.

3.  **Abra o terminal** e navegue até à pasta onde guardou o arquivo.

4.  **Execute o comando**:

    ```bash
    go run Gobattle.go
    ```

## Créditos

Desenvolvido como um projeto de demonstração de lógica de programação em Go.

**Autor**: Hudson
