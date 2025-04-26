# dotfiles
Dotfiles and ansible playbook for my windows, mac and linux workstations.

## How to use
Make sure you have ansible installed.
Just run `ansible-playbook dev.yml --ask-become-pass` and give it your password.
Thats it for linux distros and macos. For windows, install windows native gui software manually for now.

## To Do
- [X] Initial setup
- [ ] Use tags 
- [ ] Windows native apps
- [ ] Nix / Home manager integration

## Extending
- To install additional packages, just mention them in `ansible/packages.yml` under appropriate var.
- To add more dotfiles, just place them in the root directory of this repo, and declare its destination in `ansible/dots.yml`.