#!/usr/bin/env bash

echo "Uninstalling gitstart..."

# Check installation
install_path=$(command -v gitstart)

# Old and new config locations
old_config="$HOME/.gitstart_config"
new_config_dir="${XDG_CONFIG_HOME:-$HOME/.config}/gitstart"

# Remove old config if exists
if [ -f "$old_config" ]; then
    echo "Removing old config file: $old_config"
    rm "$old_config" || {
        echo "Please manually remove: $old_config"
    }
fi

# Remove new config directory if exists
if [ -d "$new_config_dir" ]; then
    echo "Removing config directory: $new_config_dir"
    rm -rf "$new_config_dir" || {
        echo "Please manually remove: $new_config_dir"
    }
fi

# Remove the script based on installation method
if [ -z "$install_path" ]; then
    echo "gitstart is not found in PATH. It may already be uninstalled."
    exit 0
fi

echo "Removing gitstart script from: $install_path"

case "$install_path" in
*local/share*)
    # awesome package manager
    echo "Detected: Awesome package manager installation"
    rm "$install_path" || {
        echo "ERROR: Failed to remove $install_path"
        echo "Please manually run: rm $install_path"
        exit 1
    }
    echo "Successfully removed gitstart"
    ;;
*brew* | */Cellar/*)
    # Homebrew installation
    echo "Detected: Homebrew installation"
    brew uninstall gitstart || {
        echo "ERROR: Failed to uninstall via Homebrew"
        echo "Please manually run: brew uninstall gitstart"
        exit 1
    }
    echo "Untapping shinokada/gitstart..."
    brew untap shinokada/gitstart 2>/dev/null || true
    echo "Successfully uninstalled gitstart"
    ;;
*/usr/bin* | */usr/local/bin*)
    # Debian package or system installation
    echo "Detected: Debian/system installation"
    if command -v apt &>/dev/null; then
        sudo apt remove gitstart -y || {
            echo "ERROR: Failed to remove via apt"
            echo "Please manually run: sudo apt remove gitstart"
            exit 1
        }
        echo "Successfully removed gitstart"
    elif command -v dpkg &>/dev/null; then
        sudo dpkg -r gitstart || {
            echo "ERROR: Failed to remove via dpkg"
            echo "Please manually run: sudo dpkg -r gitstart"
            exit 1
        }
        echo "Successfully removed gitstart"
    else
        echo "ERROR: Package manager not found"
        echo "Please manually run: sudo rm $install_path"
        exit 1
    fi
    ;;
*)
    # Unknown installation method
    echo "WARNING: Unknown installation method detected"
    echo "Installation path: $install_path"
    echo ""
    read -r -p "Do you want to remove the script manually? (y/n): " answer
    if [[ "$answer" =~ ^[Yy] ]]; then
        rm "$install_path" 2>/dev/null && echo "Successfully removed gitstart" || {
            echo "ERROR: Failed to remove $install_path"
            echo "You may need root permissions. Try: sudo rm $install_path"
            exit 1
        }
    else
        echo "Uninstallation cancelled"
        echo "To manually uninstall, run: rm $install_path"
        exit 0
    fi
    ;;
esac

echo ""
echo "===================================="
echo "Gitstart uninstallation completed!"
echo "===================================="
echo ""
echo "Removed:"
echo "  - Script: $install_path"
if [ -f "$old_config" ] || [ -d "$new_config_dir" ]; then
    echo "  - Configuration files"
fi
echo ""
echo "Thank you for using gitstart!"
