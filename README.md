# Torneio suíço

Este projeto simula um torneio de jogadores, onde os participantes competem em várias rodadas. O sistema organiza os jogadores de acordo com o número de vitórias, gerencia empates, byes (quando um jogador não tem oponente), e calcula a classificação final levando em consideração a pontuação de Buchholz.

# Funcionalidades
Emparelhamento de jogadores: Os jogadores são emparelhados aleatoriamente, com base no número de vitórias. Caso o número de jogadores seja ímpar, um jogador ganha um "bye", ou seja, avança sem competir naquela rodada.
Registro de resultados: O sistema permite registrar o vencedor de cada partida, atualizando automaticamente a classificação dos jogadores.
Cálculo de Buchholz: Após cada rodada, o sistema calcula a pontuação de Buchholz, que é usada para desempatar jogadores com o mesmo número de vitórias.
Classificação final: Ao final do torneio, o sistema exibe a classificação final dos jogadores, considerando vitórias, direct wins, e a pontuação de Buchholz.
Como Usar
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

Estrutura de Dados
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

Exemplo de Uso
bash
Copiar
Editar
Round 1:
Match 1: Anthony vs Renata
Match 2: Henrique vs Sofia
Match 3: Thiago vs Arthur
Match 4: Renato gets a bye.

Enter result for Match 1: Anthony vs Renata
Enter result (1 for Player1, 2 for Player2, 0 for draw): 1

...

Final Standings After Recalculating Buchholz Scores:
ID: 1, Name: Anthony, Wins: 3, Byes: 0, Buchholz: 5
ID: 2, Name: Renata, Wins: 2, Byes: 1, Buchholz: 3
...
Licença
Este projeto está licenciado sob a Licença MIT.