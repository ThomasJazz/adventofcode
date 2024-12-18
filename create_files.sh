# parse args and create files and directories based on the args
# I am no bash god, this was all AI generated lol

POSITIONAL_ARGS=()
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        -y|--year)
            year="$2"
            shift
            shift
            ;;
        -d|--day)
            day="$2"
            shift
            shift
            ;;
        *)
            POSITIONAL_ARGS+=("$1")
            shift
            ;;
    esac
done

set -- "${POSITIONAL_ARGS[@]}"
if [[ -z "$year" || -z "$day" ]]; then
    echo "Usage: $0 -y <year> -d <day>"
    exit 1
fi

dir_path="$(pwd)/$year/solutions/day$day"
mkdir -p "$dir_path"

cat <<EOL > "$dir_path/day$day.go"
package day$day

import (
    "fmt"
)

var (
    filename = "../../input/input-day${day}.txt"
)

func PartOne() {
    fmt.Println("Day $day solution pt 1")
}

func PartTwo() {
    fmt.Println("Day $day solution pt 2")
}
EOL

cat <<EOL > "$dir_path/day${day}_test.go"
package day$day

import (
    "testing"
)

func TestPartOne(t *testing.T) {
    // Add tests here
}

func TestPartTwo(t *testing.T) {
    // Add tests here
}
EOL

input_path="$(pwd)/$year/input"
mkdir -p "$input_path"
touch "$input_path/input-day$day.txt"

echo "Created $dir_path/day$day.go and $dir_path/day$day_test.go and $input_path/input-day$day.txt" 

