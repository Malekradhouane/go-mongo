#!/bin/bash

# Run your changelog generator script and save the output to a file
changelog > changelog.txt

# Read the changelog content from the file
CHANGELOG_CONTENT=$(cat changelog.txt)

# Set the message template with the changelog content
MESSAGE_TEMPLATE="iTower $1 is out!\n\nChangelog:\n$CHANGELOG_CONTENT"

# Output the message template
echo "$MESSAGE_TEMPLATE"