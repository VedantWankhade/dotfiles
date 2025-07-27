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

> [!IMPORTANT]  
> Use ansible-core v2.18.*
> v2.19.* has issues with reading sudo password from tty stdin as of 27/07/2025.

> [!IMPORTANT]
> Turns out there are some breaking changes with ansible-core v2.19.0. 
> > Task Execution / Forks - Forks no longer inherit stdio from the parent ansible-playbook process. stdout, stderr, and stdin within a worker are detached from the terminal, and non-functional. All needs to access stdio from a fork for controller side plugins requires use of Display - 
> https://github.com/ansible/ansible/blob/v2.19.0/changelogs/CHANGELOG-v2.19.rst#major-changes