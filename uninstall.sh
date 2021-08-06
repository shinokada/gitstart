#!/usr/bin/env bash

echo "Uninstalling gitstart..."

# check installation
install_path=$(command -v gitstart)
gitstart_config=$HOME/.gitstart_config

# awesome installation path $HOME/.local/share/bin
# brew installation path $(brew --prefix)/bin
# Ubuntu installtion path /usr

if [ -f "$gitstart_config" ]; then
    echo "Removing gitstart_config..."
    rm "$gitstart_config" || {
        echo "Please removed $gitstart_config."
    }
fi

echo "Removing gitstart script..."

case "$install_path" in
*local/share*)
    # awesome
    rm "$install_path" || {
        echo "Please remove $install_path."
    }
    ;;
*brew*)
    # brew
    brew uninstall gitstart || {
        echo "Please remove $install_path."
    }
    brew untap shinokada/gitstart
    ;;
*usr/bin*)
    # debian package
    sudo apt remove gitstart || {
        echo "Please remove $install_path."
    }
    ;;
*)
    # unknown
    echo "Not able to find your installation method."
    echo "Please uninstall gitstart script."
    ;;
esac

echo "Uninstalltion completed."
