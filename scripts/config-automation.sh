#!/bin/bash

SHELL_CONFIG=""
if [ -f "$HOME/.zshrc" ]; then
    SHELL_CONFIG="$HOME/.zshrc"
elif [ -f "$HOME/.bashrc" ]; then
    SHELL_CONFIG="$HOME/.bashrc"
else
    SHELL_CONFIG="$HOME/.profile"
fi

START="sugestion-start"
END="sugestion-end"

if grep -q "$START" "$SHELL_CONFIG" || grep -q "git-commit-tool" "$SHELL_CONFIG"; then
    echo "Existing sugestion found in $SHELL_CONFIG. Removing it first."

    sed -i '' '/git-commit-tool/d' "$SHELL_CONFIG" 2>/dev/null || true
    sed -i '' "/$START/,/$END/d" "$SHELL_CONFIG" 2>/dev/null || true
fi

echo "Installing sugestion in $SHELL_CONFIG"

cat << EOF >> "$SHELL_CONFIG"

# $START
git() {
    if [ "\$1" = "add" ]; then
        command git "\$@"
        if [ \$? -eq 0 ]; then
            make run
        fi
    else
        command git "\$@"
    fi
}
# $END

EOF

echo "Installation complete. Please restart your terminal or run 'source $SHELL_CONFIG' to apply the changes."
