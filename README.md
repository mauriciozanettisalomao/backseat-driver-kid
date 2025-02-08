# Backseat Driver Kid 🚗👦

A CLI tool that’s like having a **curious toddler**

## Personal Purpose

The main goal of this project is to learn how to work with technologies that allow data analysis without exposing sensitive information to the internet (e.g., through APIs). The project also aims to:

* Learn more about Large Language Models (LLMs) and the libraries that support them.
* Understand how Retrieval-Augmented Generation (RAGs), vectors, and databases work together.
* Explore best practices for using LLMs, such as designing prompts and generating responses.

## Introduction

**Kid:** Dad, why do we have to read all these files again? Can't we just play video games instead?

**Dad:** Because, kiddo, this project is about learning how to work with Large Language Models (LLMs) like Llama2 using the Ollama framework and the LangChainGo library. If you want to build smart applications, you need to understand how they work.

**Kid:** But why do we have to read all these files? Can't we just use the magic button thingy again?

**Dad:** The "magic button thingy"? That’s not how coding works, buddy. You have to install the right dependencies and set up the project properly. Here, let’s start with the installation.

---

## Installation

1. Clone the repository:

   ```sh
   git clone github.com/mauriciozanettisalomao/backseat-driver-kid.git
   ```

2. Navigate to the project directory:

   ```sh
   cd backseat-driver-kid
   ```

3. Install dependencies:

   ```sh
   go mod tidy
   ```

4. Install Ollama to run LLM locally

   ```sh
    curl -fsSL https://ollama.ai/install.sh | sh
   ```

5. Install the application:

   ```sh
   go build
   ```

6. Run the application:

   ```sh
   ./backseat-driver-kid prompt apply 
   ```

**Kid:** Okay, so we type these commands into the terminal, and then it works?

**Dad:** Exactly! But before running it, we need to ensure that the framework and our **LLM models** are loaded.

   ```sh
    ollama --version
   ```

_ollama run {{model}}_ - e.g.:

   ```sh
    ollama run llama2
   ```

## 📌 Usage

### **Basic Help Command**

Get an overview of available commands:

```sh
backseat-driver-kid --help
```

### **Prompt Commands**

To interact with the CLI and apply a prompt:

    ```sh
    backseat-driver-kid prompt apply --input resources/input/interaction-config.yaml
    ```

#### **Flags**

| Flag         | Description                                   | Default                                   |
| ------------ | --------------------------------------------- | ----------------------------------------- |
| `--input`    | YAML config file for interaction              | `resources/input/interaction-config.yaml` |
| `--model`    | LLM model to use (e.g., Llama2)               | `llama2`                                  |
| `--output`   | Output file for the generated responses       | `resources/output/analysis.md`            |
| `--routines` | Number of goroutines for concurrent execution | `1`                                       |

Example with custom values:

    ```sh
    backseat-driver-kid prompt apply --input myconfig.yaml --model mistral --output output.md --routines 5
    ```

## 📝 Configuration File Format

The tool expects a **YAML** configuration file structured as follows:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: interaction-config
data:
  interaction:
    preamble:
      context: This is a simulated interaction between a user and a system.
      instructions: Provide clear and concise answers based on the context provided.
      examples: |
        Example 1: User asks about system health, and the response should indicate the status.
        Example 2: User asks for an operation status, and the system should provide a detailed report.
    # The path to the directory containing the knowledge files in text format
    # that contain the information to be used in the interaction.
    extendedKownledgeDir:
      - resources/input/knowledge
    ## list of prompts that the user can ask
    prompts:
      - input: What is the status of the system?
      - input: How many incidents have been reported?
    ## there's an option to provide a file with a list of prompts
    promptFile: resources/input/prompts/questions.txt
```

## 📡 Logging & Debugging

Enable debugging by setting the log level:
    ```sh
    backseat-driver-kid prompt apply --log-level debug
    ```

Change log format to plain text:
    ```sh
    backseat-driver-kid prompt apply --log-format text
    ```

## 🤔 Need Help?

Check available commands:
    ```sh
    backseat-driver-kid help
    ```

For a specific command:
    ```sh
    backseat-driver-kid prompt apply --help
    ```

--- 

## Features
- **Integration with Ollama and Llama2** for advanced language processing.
- **Secure handling of personal data** to ensure privacy.
- **Utilization of RAG (Retrieval-Augmented Generation)** 🛠️ Not implemented yet 🚧

**Kid:** RAG? Vectors? That sounds complicated. Can we watch a movie instead?

**Dad:** Not yet, but don’t worry, it’ll be fun once you understand it! Think of **vectors** as coordinates that help us find the best answer, and **RAG** as a way to combine the best answers from multiple sources. It makes the AI smarter!

---

## Security Considerations

**Kid:** This is getting long. Why does security matter?

**Dad:** Because we’re working with **personal data**, and it’s crucial to **never expose sensitive information**. We follow these best practices:

**Kid:** Okay… I guess protecting data is important. But can we get ice cream if we finish?

**Dad:** Deal! But first, let’s check out some useful resources.

---

## Learning Resources

- [Ollama Documentation](https://github.com/ollama/ollama)
- [Llama2 Model Overview](https://ollama.com/library/llama2)
- [LangChain Go](https://github.com/tmc/langchaingo)

## To Do

* *RAG Implementation:* Adding support for Retrieval-Augmented Generation (RAG).
* *Unit Testing:* Writing and implementing unit tests for the core functionality.
* *Makefile:* Creating a Makefile for easier project setup and management.
* *Linter:* Adding a linter to enforce code style and standards.
* *Security Scans:* Implementing security scans to ensure code safety and security.
* *Container Support:* Adding support for Docker containers to run the application in isolated environments.
* *GitHub Actions:* Setting up GitHub Actions for pull request checks, releases, and other CI/CD tasks.

## Contributing
Pull requests are welcome! Follow the **Go project structure** and best practices when submitting changes.

## License
This project is licensed under the MIT License.

**Kid:** Finally done! Now, ice cream?

**Dad:** Yes, but remember—reading the docs first makes everything easier!
