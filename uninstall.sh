#!/usr/bin/env bash

echo "Uninstalling gitstart..."

# check installation
install_path=$(command -v gitstart)
gitstart_config=$HOME/.gitstart_config

echo "Removing gitstart_config..."
rm "$gitstart_config" || {
    echo "Please removed $gitstart_config."
}

echo "Removing gitstart script..."
rm "$install_path" || {
    echo "Please removed $install_path."
}

echo "Uninstalltion completed."
