package main

import (
	"fmt"
	"math/rand"
	"time"
)


type Personagem interface {
	GetNome() string
	GetVida() int
	Atacar() int
	ReceberDano(dano int)
}


// GUERREIRO
type Guerreiro struct {
	Nome string
	Vida int
}

func (g Guerreiro) GetNome() string { return g.Nome }
func (g Guerreiro) GetVida() int   { return g.Vida }
func (g *Guerreiro) ReceberDano(dano int) {
	g.Vida -= dano
	if g.Vida < 0 {
		g.Vida = 0
	}
}
func (g Guerreiro) Atacar() int {
	dano := rand.Intn(5) + 3 
	fmt.Printf("⚔️  %s deu uma espadada e causou %d de dano!\n", g.Nome, dano)
	return dano
}

// MAGO 
type Mago struct {
	Nome string
	Vida int
}

func (m Mago) GetNome() string { return m.Nome }
func (m Mago) GetVida() int   { return m.Vida }
func (m *Mago) ReceberDano(dano int) {
	m.Vida -= dano
	if m.Vida < 0 {
		m.Vida = 0
	}
}
func (m Mago) Atacar() int {
	dano := rand.Intn(8) + 2 
	fmt.Printf("%s lançou uma bola de fogo e causou %d de dano!\n", m.Nome, dano)
	return dano
}

// MONSTRO 
type Monstro struct {
	Vida int
}

func (m Monstro) Atacar() int {
	return rand.Intn(5) + 1 
}

func IniciarCombate(jogador Personagem, monstro *Monstro) {
	fmt.Printf("\n--- O COMBATE COMEÇOU! O Monstro tem %d de vida ---\n", monstro.Vida)

		for jogador.GetVida() > 0 && monstro.Vida > 0 {
		danoDoJogador := jogador.Atacar()
		monstro.Vida -= danoDoJogador
		fmt.Printf("Vida do Monstro: %d\n", monstro.Vida)

		if monstro.Vida <= 0 {
			fmt.Println("\n Você derrotou o monstro!")
			return
		}

		danoDoMonstro := monstro.Atacar()
		jogador.ReceberDano(danoDoMonstro)
		fmt.Printf("O Monstro atacou! Causou %d de dano.\n", danoDoMonstro)
		fmt.Printf("Sua Vida: %d\n\n", jogador.GetVida())

		if jogador.GetVida() <= 0 {
			fmt.Println("\n VOCÊ MORREU! FIM.")
			return
		}

		time.Sleep(1 * time.Second)
	}
}

func main() {
	var escolha int
	var nome string

	fmt.Println("Bem-vindo ao Go RPG!")
	fmt.Print("Digite o nome do seu herói: ")
	fmt.Scan(&nome)

	fmt.Println("\nEscolha sua classe:")
	fmt.Println("1) Guerreiro (Vida: 20)")
	fmt.Println("2) Mago (Vida: 10)")
	fmt.Print("Digite o número: ")
	fmt.Scan(&escolha)

	var jogador Personagem
	if escolha == 1 {
		jogador = &Guerreiro{Nome: nome, Vida: 20}
	} else if escolha == 2 {
		jogador = &Mago{Nome: nome, Vida: 10}
	} else {
		fmt.Println("Escolha inválida! Você vai de Guerreiro por padrão.")
		jogador = &Guerreiro{Nome: nome, Vida: 20}
	}

	vidaDoMonstro := rand.Intn(16) + 5 
	chefao := &Monstro{Vida: vidaDoMonstro}

	IniciarCombate(jogador, chefao)
}