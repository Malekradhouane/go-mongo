#!/bin/bash

branch_name=$(echo $CI_MERGE_REQUEST_SOURCE_BRANCH_NAME | cut -d'/' -f 1)

echo $branch_name

x=$(git describe --tags $(git rev-list --tags --max-count=1))

echo $x


if [[ $branch_name == fix/* ]]; then
    version=$(git describe --tags `git rev-parse HEAD`)
    release=$(echo $version | cut -d. -f1)
    feature=$(echo $version | cut -d. -f2)
    fix=$(echo $version | cut -d. -f3)
    new_fix=$((fix + 1))
    new_version="$release.$feature.$new_fix"
elif [[ $branch_name == feat/* ]]; then
        version=$(git describe --tags `git rev-parse HEAD`)
        release=$(echo $version | cut -d. -f1)
        feature=$(echo $version | cut -d. -f2)
        fix=$(echo $version | cut -d. -f3)
        new_feature=$((feature + 1))
        new_version="$release.$feature.$new_feature"
elif [[ $branch_name == release/* ]]; then
         version=$(git describe --tags `git rev-parse HEAD`)
         release=$(echo $version | cut -d. -f1)
         feature=$(echo $version | cut -d. -f2)
         fix=$(echo $version | cut -d. -f3)
         new_release=$((release + 1))
         new_version="$release.$feature.$new_release"
fi
echo $new_version
echo $version
