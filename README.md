# 🐉 Pokedex CLI

A command-line interface tool built in Go to interact with a simulated Pokémon API, allowing users to explore locations, search for Pokémon, and manage a local Pokedex inventory.

This project simulates the functionality of a comprehensive Pokémon database utility, providing an interactive Read-Eval-Print Loop (REPL) experience directly in your terminal.

## ✨ Features

*   **Interactive REPL:** Use a simple command structure to interact with the Pokedex without writing full programs.
*   **Location Exploration:** Fetch Pokémon lists for specific areas using `explore`.
*   **Inventory Management:** View all caught Pokémon using `pokedex`.
*   **Data Inspection:** Get detailed information about a specific Pokémon using `inspect`.
*   **Batch Fetching:** Fetch lists of sequential locations (previous/next areas) using `mapf` and `mapb`.

## 🚀 Getting Started

### Prerequisites

You need the following installed on your system:
*   Go (Golang)
*   A basic understanding of CLI tools.

### Installation

1.  **Clone the Repository:**
    ```bash
    git clone github.com/valbertoenoc/pokedexcli
    cd pokedex
    ```

2.  **Install Dependencies:**
    The project relies on internal modules, which should be initialized:
    ```bash
    # Ensure you are in the root directory (pokedex/)
    go mod tidy
    ```

3.  **Run the Application:**
    Execute the main file to start the interactive shell:
    ```bash
    go run main.go
    ```

## 🧭 Usage Guide (The REPL)

Once the application starts, you will see the `Pokedex >` prompt. Type a command and press Enter.

### Command Reference

| Command | Usage | Description |
| :--- | :--- | :--- |
| `help` | `help` | Displays this help message and command options. |
| `explore` | `explore <location_name>` | Fetches and lists all Pokémon found in a specified area. |
| `inspect` | `inspect <pokemon_name>` | Retrieves and prints detailed statistics for a given Pokémon. |
| `catch` | `catch <pokemon_name>` | Simulates attempting to capture a Pokémon and adds it to your local Pokedex. |
| `pokedex` | `pokedex` | Displays a list of all Pokémon currently stored in your Pokedex inventory. |
| `mapf` | `mapf` | Fetches a batch of 20 next sequential location areas. |
| `mapb` | `mapb` | Fetches a batch of 20 previous sequential location areas. |
| `exit` | `exit` | Exits the interactive Pokedex tool. |

**Example Flow:**
```
Pokedex > help
Pokedex > explore mt-coronet-3f
... (Lists Pokémon found in Tall Grass)
Pokedex > catch pikachu
Success! Pikachu added to your Pokedex.
Pokedex > pokedex
Pokedex: [Pikachu]
```

## 🛠️ Architecture Overview

The CLI logic is structured to provide a clean, command-dispatching pattern:

*   **`main.go`**: Initializes the client and starts the REPL loop.
*   **`repl.go`**: Contains the core `startRepl` function, which reads user input, cleans it, and dispatches the command to the appropriate callback.
*   **`internal/pokeapi`**: Handles all external API communication and data fetching logic.
*   **`internal/pokecache`**: Handles application level cache logic.
*   **`config` struct**: Manages the state of the application (API client, captured Pokémon cache, etc.) passed between commands.
