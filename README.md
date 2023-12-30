# Gazes Novels: Guidelines for Contributions
*Crafted with ❤️ by l31*

Greetings, esteemed developers! Welcome to the vibrant realm of **Gazes Novels**, a dynamic hub dedicated to the creation and enjoyment of web novels. If you're eager to contribute your coding expertise to this noble endeavor, please acquaint yourself with the rules and guidelines outlined in this comprehensive document.

## Project Overview

**Gazes Novels** is ingeniously crafted using the Go programming language, with GORM serving as our trusted Object-Relational Mapping (ORM) tool and PostgreSQL as the preferred database. It's important to note that our tech stack is a dynamic entity, open to evolution, and this document will gracefully adapt to any changes.

## Making Your Impact

### Setting the Foundation

Before immersing yourself in the thrilling world of contribution, ensure you have [nix](https://nixos.org/) installed. Fear not, as we've streamlined the process with a one-liner, courtesy of [zero-to-nix](https://zero-to-nix.com/). Execute the following command to initiate the installation:

```bash
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
```

Follow the steps presented by the installer. Once completed, a touch of configuration awaits. Activate some experimental features, specifically `nix-command` and `nix-flakes`. Edit either the file `/etc/nix/nix.conf` or `~/.config/nix/nix.conf` and append the following line:

```conf
experimental-features = nix-command flakes
```

Save your changes, and voilà! You're all set for the next step!

### Embark on Your Coding Journey

Fantastic! With nix securely installed, let's delve into the exhilarating realm of coding. Begin your journey by forking, cloning, and opening a development shell. It's worth noting that we maintain a dedicated shell for each project within this mono-repo.

1. **Fork & Clone:**

    Fork the project and proceed to clone it using the following command:

    ```bash
    git clone https://github.com/<your-username>/gazes-novels.git
    ```

2. **Navigate to Project:**

    Move into the project you're eager to contribute to; for instance, let's explore the `backend` project:

    ```bash
    cd gazes-novels/backend
    ```

3. **Launch the Development Shell:**

    Take the next step by opening the development shell:

    ```bash
    nix develop
    ```

    Well done! Your coding environment is ready. Let the coding festivities commence!