# CHIP_8_emu
emulator for chip 8
Ce code est un programme en Go qui implémente un émulateur de la console de jeu CHIP-8. Je vais expliquer chaque partie du code en détail.

Tout d'abord, le code commence par importer plusieurs packages nécessaires :

	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"os"

fmt est le package standard pour les fonctions d'entrée/sortie.
github.com/hajimehoshi/ebiten/v2 est une bibliothèque qui permet de créer des jeux en 2D en utilisant Ebiten, un moteur de jeu en Go.
github.com/hajimehoshi/ebiten/v2/ebitenutil est un sous-package d'Ebiten qui fournit des utilitaires pour les jeux.
image/color est utilisé pour la gestion des couleurs.
log est utilisé pour la journalisation.
os est utilisé pour les opérations sur le système d'exploitation.
Ensuite, le code définit plusieurs structures de données pour représenter l'état de l'émulateur CHIP-8. Ces structures comprennent CPU, Chip8, Clavier, et Screen. Voici un aperçu de chacune :

CPU représente le processeur CHIP-8 avec des champs pour la mémoire, les registres, le compteur de programme (PC), l'indice (I), la pile (stack), etc.

Chip8 est une structure qui contient le CPU, un clavier (qui n'est pas encore implémenté), et un écran.

Clavier est destiné à représenter les entrées du clavier, mais dans le code actuel, il n'est pas utilisé.

Screen représente l'écran du CHIP-8 avec une matrice de pixels de 64x32.

Ensuite, il y a des variables globales déclarées, notamment chip8, ROM, et screen.

chip8 est une instance de la structure Chip8 qui représente l'émulateur CHIP-8.

ROM est une variable qui stockera le contenu du fichier ROM (le programme à exécuter).

screen est une instance de l'écran Ebiten qui sera utilisée pour afficher la sortie graphique du jeu.

Le code définit également quelques constantes, telles que screenWidth, screenHeight, resolWidth, et resolHeight, qui définissent les dimensions de l'écran.

Ensuite, le code définit la structure Console qui représente l'état du jeu CHIP-8. Cette structure inclut les champs IN (entrée), OUT (sortie), et command (commande). Cependant, dans le code actuel, ces champs ne sont pas utilisés.

La fonction Init est utilisée pour initialiser certaines parties de l'émulateur CHIP-8. Elle remplit la mémoire du CHIP-8 avec un ensemble de caractères spécifiques (fontSet), qui sont les caractères affichables par le CHIP-8.

La fonction NewConsole crée une nouvelle instance de la structure Console.

Les méthodes Update, Draw, et Layout de la structure Console sont utilisées pour mettre à jour l'état du jeu, dessiner l'écran, et gérer la mise en page de l'écran, respectivement. Cependant, dans le code actuel, elles sont incomplètes et ne font pas grand-chose.

La fonction LoadProgram configure la fenêtre Ebiten, crée une instance de Console, définit le nombre de trames par seconde (TPS), et exécute le jeu à l'aide de la fonction ebiten.RunGame.

La fonction LoadROM est utilisée pour charger un programme ROM dans la mémoire du CHIP-8.

La fonction Interpreter est le cœur de l'émulateur CHIP-8. Elle interprète les instructions du CHIP-8 à partir de la mémoire et met à jour l'état en conséquence. Elle contient une grande structure de commutation pour gérer différents types d'instructions du CHIP-8.

La fonction Start est le point de départ de l'émulateur. Elle lit un fichier ROM, initialise l'émulateur, charge le ROM dans la mémoire, configure l'écran, et exécute le jeu.

Enfin, la fonction main appelle Start pour démarrer l'émulateur CHIP-8.

En résumé, ce code est le début d'un émulateur CHIP-8 en Go. Il définit les structures de données et les fonctions nécessaires pour exécuter des programmes CHIP-8, mais de nombreuses parties, telles que la gestion de l'entrée du clavier et la mise à jour de l'écran, ne sont pas encore complètement implémentées. L'émulateur est également conçu pour utiliser la bibliothèque Ebiten pour la gestion de la fenêtre et de l'affichage graphique.
