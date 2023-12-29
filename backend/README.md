# Setting Up Your Development Environment

Before you start writing code for this project, it's crucial to establish the right development environment. Follow these steps to ensure a smooth setup:

## Step 1: Install Nix

If you don't have Nix installed on your system, you need to install it first. Follow the instructions in [this guide](https://zero-to-nix.com/start/install) to set up Nix. For Windows users, makes sure to enable and use [WSL2](https://learn.microsoft.com/fr-fr/windows/wsl/install).

## Step 2: Activate Nix Command and Flake Experimental Features

If you haven't already activated the Nix command and flake experimental features, add the following line to your `~/.config/nix/nix.conf` file:

```bash
experimental-features = nix-command flakes
```

This step enables additional features that enhance the functionality of Nix, providing you with a more powerful development experience.

## Step 3: Initialize Backend Development Environment

Once Nix is set up, initiate the development environment specifically tailored for backend work in this project. Run the following command:

```bash
nix develop .#backend
```

This command ensures that the necessary dependencies and tools for backend development are set up in your environment. Execute this step to kickstart your coding journey seamlessly.

By following these steps, you'll have a well-configured development environment ready for coding in the project. Happy coding!