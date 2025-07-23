# dotfiles
Dotfiles and ansible playbook for my windows, mac and linux workstations (and my VPS??).

## How to use
### With Ansible
- Make sure you have ansible installed.
- Just run `ansible-playbook dev.yml --ask-become-pass` and give it your password.
- Thats it for linux distros and macos. For windows, install windows native gui software manually for now.
### With Go
- Just run `go run github.com/vedantwankhade/dotfiles@latest`

## Extending
- To install additional packages, just mention them in `<os>/packages.yml` under appropriate var.
- To add more dotfiles, just place them in the root directory of this repo, and declare its destination in `<os>/dots.yml`.

## To Do
- OS choice (like WSL, Arch etc)
- Profile choices (like Hyprland, DWM, KDE, Dev etc)
- Hyprland
  - Switch to specific window (or open the program if not running) on keybind