import random

class ChuteONumero:
    def __init__(self):
        self.valor_aleatorio = 0
        self.valor_minimo = 1
        self.valor_maximo = 100
        self.total_de_tentativas = 0
        self.pontos = 1000

    def Intro(self):
        print("        Bem-vindo ao...       ")
        print("____ JOGO DA ADIVINHAÇÃO! ____")
        print(" Descubra o número de 0 a 100" )

    def Dificuldade(self):
        print("          Dificuldades           ")
        print("(1) Fácil  (2) Médio  (3) Difícil")
        nivel = int(input("Escolha: "))
        if (nivel == 1):
            self.total_de_tentativas = 15
            print("Começando com calma...")
        elif (nivel == 2):
            self.total_de_tentativas = 10
            print("Equilíbrio é tudo!")
        else:
            self.total_de_tentativas = 5
            print("Ai, ignorante.")

    def Iniciar(self):
        for rodada in range(1, self.total_de_tentativas + 1):
            print("Tentativa {} de {}".format(rodada, self.total_de_tentativas))
            self.PedirValorAleatorio()
            if int(self.valor_do_chute) == self.valor_aleatorio:
                print("PARABÉNS, VOCÊ ACERTOU!")
                print("E fez {} pontos".format(self.pontos))
                break
            else:
                if int(self.valor_do_chute) > self.valor_aleatorio:
                    print("Um pouco mais pra baixo")
                elif int(self.valor_do_chute) < self.valor_aleatorio:
                    print("Um pouco mais pra cima")

                self.pontos_perdidos = abs(self.valor_aleatorio - int(self.valor_do_chute))
                self.pontos = self.pontos - self.pontos_perdidos

    def PedirValorAleatorio(self):
        self.valor_do_chute = input("Chute um número: ")

    def GerarNumeroAleatorio(self):
        self.valor_aleatorio = random.randint(self.valor_minimo, self.valor_maximo)



chute = ChuteONumero()
chute.GerarNumeroAleatorio()
chute.Intro()
chute.Dificuldade()
chute.Iniciar()