# Torneio suíço

Este projeto simula um torneio de jogadores usando o formato Torneio Suíço. Os jogadores competem em várias rodadas, sendo emparelhados com base em suas vitórias anteriores. O sistema também gerencia empates, byes (quando um jogador não tem oponente), e calcula a classificação final considerando a pontuação de Buchholz para desempate.

# Funcionalidades

1. Emparelhamento de jogadores: Os jogadores são emparelhados de forma que os jogadores com o mesmo número de vitórias enfrentem-se, mas com aleatoriedade dentro do mesmo grupo de vitórias. Caso o número de jogadores seja ímpar, um jogador ganha um "bye", ou seja, avança sem competir naquela rodada.
2. Registro de resultados: O sistema permite registrar o vencedor de cada partida, atualizando automaticamente a classificação dos jogadores.
3. Cálculo de Buchholz: Após cada rodada, o sistema calcula a pontuação de Buchholz, que é usada para desempatar jogadores com o mesmo número de vitórias.
4. Classificação final: Ao final do torneio, o sistema exibe a classificação final dos jogadores, considerando vitórias, direct wins, e a pontuação de Buchholz.

# Como Funciona o Torneio Suíço
O formato Torneio Suíço é utilizado para competições onde todos os jogadores participam até o final do torneio, mesmo que percam uma ou mais partidas. Ao contrário de um torneio de eliminação direta, onde um jogador é eliminado a cada derrota, no torneio suíço os jogadores continuam competindo até o final, enfrentando adversários com um número similar de vitórias.

## Emparelhamento
O emparelhamento de jogadores em cada rodada é feito com base no número de vitórias acumuladas até o momento. Quando os jogadores têm o mesmo número de vitórias, o sorteio aleatório é realizado dentro desse grupo de jogadores. Caso o número de jogadores seja ímpar, um jogador é sorteado para receber um bye, ou seja, ele não competirá naquela rodada e será considerado como vencedor, ganhando uma vitória sem jogar.

## Sorteio
Os jogadores são embaralhados antes do torneio começar, e as partidas são sorteadas de acordo com o número de vitórias. Quando os jogadores têm o mesmo número de vitórias, um sorteio aleatório é realizado dentro desse grupo para decidir os emparelhamentos.

# Critérios de Desempate
Quando dois ou mais jogadores têm o mesmo número de vitórias, o desempate é realizado de acordo com os seguintes critérios:

1.  Vitórias Diretas: Caso dois jogadores empatem em vitórias, o critério de desempate é o histórico de vitórias diretas entre eles. O jogador que venceu o outro em uma rodada anterior ganha a posição mais alta.

2. Buchholz: Se os jogadores ainda estiverem empatados após o critério de vitórias diretas, a pontuação de Buchholz é usada. A pontuação de Buchholz é calculada somando o número de vitórias dos adversários vencidos por cada jogador. Quanto maior a pontuação de Buchholz, melhor o desempenho do jogador.

3. ID do Jogador: Se ainda houver empate, o critério final é o ID do jogador, com jogadores de ID menor ficando à frente.

# Como Usar
Configuração Inicial:

O torneio começa com a criação de um novo objeto Tournament, onde você fornece os nomes dos jogadores.
O sistema organiza os jogadores e inicia as rodadas.
Rodadas e Emparelhamentos:

A cada rodada, os jogadores são emparelhados aleatoriamente, com exceção dos que receberam um bye.
O emparelhamento de jogadores é mostrado antes de pedir a entrada dos resultados de cada partida.
Entrada de Resultados:

O sistema solicita a entrada do resultado de cada partida, informando claramente qual é o match que você está registrando.
O formato esperado para a entrada é:
1 para vitória do jogador 1
2 para vitória do jogador 2
0 para empate
Cálculo e Exibição da Classificação:

Após cada rodada, a classificação é atualizada e mostrada na tela, considerando vitórias, empates e a pontuação de Buchholz.
A classificação final é exibida ao término do torneio.
Como Rodar
Para rodar o programa, basta compilar e executar o código Go.
Você poderá seguir as instruções para inserir os resultados das partidas e acompanhar a classificação dos jogadores.

# Estrutura de Dados
Player: Representa um jogador, com os campos:

ID: Identificador único do jogador.
Name: Nome do jogador.
Wins: Número de vitórias.
Byes: Número de vezes que o jogador recebeu um bye.
DirectWins: Um mapa que registra quais jogadores foram derrotados diretamente por este jogador.
Buchholz: A pontuação de Buchholz do jogador.
Match: Representa uma partida entre dois jogadores.

Player1: O primeiro jogador.
Player2: O segundo jogador.
Winner: Indica o vencedor da partida (1 para o jogador 1, 2 para o jogador 2, 0 para empate).
IsBye: Indica se o jogador teve um bye.
Tournament: Representa o torneio, contendo a lista de jogadores e as rodadas do torneio.