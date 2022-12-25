package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// folders are represented by a map of name to map of subfolders
type Folder struct {
	name       string
	parent     *Folder
	subfolders map[string]*Folder
	size       int
}

func newFolder(name string, parent *Folder) *Folder {
	f := Folder{
		name:       name,
		parent:     parent,
		subfolders: make(map[string]*Folder),
		size:       0,
	}
	return &f
}

var total_size = 0
var folder_sizes = []int{}

// starting at root, recursively add size of all subfolders to their parent folders
// if the size is less than 100000 add it to total_size
func addSizes(folder *Folder) int {
	size := folder.size
	for _, subfolder := range folder.subfolders {
		subfolder_size := addSizes(subfolder)
		folder_sizes = append(folder_sizes, subfolder_size)
		size += subfolder_size
		if subfolder_size < 100000 {
			total_size += subfolder_size
		}
	}
	folder.size = size
	//fmt.Printf("%+v  -  %d\n", folder, total_size)
	return size
}

func part_1(input string) (string, error) {
	// if line starts with $, it is a command '$ ls'
	// folders start with dir 'dir a'
	// files start with a number followed by their name '123 b.txt'
	// ls lists the contents of the current folder
	// cd changes the current folder, '$ cd ..' goes up one folder

	// split lines by newline
	lines := strings.Split(input, "\n")

	// iterate through lines, building a map of folders
	root := newFolder("/", nil)
	current_folder := root

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		// split line by spaces
		fields := strings.Fields(line)

		// check if line is a command
		if fields[0] == "$" {
			// ls can be skipped
			// check if command is cd
			if fields[1] == "cd" {
				// if 'cd ..' go up one folder (parent)
				if fields[2] == ".." {
					current_folder = current_folder.parent
				} else {
					// check if folder exists in all_folders
					folder, ok := current_folder.subfolders[fields[2]]
					if ok {
						// if folder exists, change current_folder to it
						current_folder = folder
					} else {
						// if folder doesn't exist, create it, switch to it
						new_folder := newFolder(fields[2], current_folder)
						current_folder.subfolders[fields[2]] = new_folder
						current_folder = new_folder
					}

				}
			}
		} else { // line is a folder or file from ls command
			// ignore any lines that start with dir since they will be
			// entered as a subfolder
			if fields[0] != "dir" {
				// parse first number as int and add it to the current folder's size
				size, err := strconv.Atoi(fields[0])
				if err != nil {
					return "", err
				}
				current_folder.size += size
			}
		}
		//fmt.Printf("%s  -  %+v, %p\n", line, current_folder, current_folder)
	}
	total_size = 0
	addSizes(root)

	return fmt.Sprint(total_size), nil
}

func part_2(input string) (string, error) {
	// if line starts with $, it is a command '$ ls'
	// folders start with dir 'dir a'
	// files start with a number followed by their name '123 b.txt'
	// ls lists the contents of the current folder
	// cd changes the current folder, '$ cd ..' goes up one folder

	// split lines by newline
	lines := strings.Split(input, "\n")

	// iterate through lines, building a map of folders
	root := newFolder("/", nil)
	current_folder := root

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		// split line by spaces
		fields := strings.Fields(line)

		// check if line is a command
		if fields[0] == "$" {
			// ls can be skipped
			// check if command is cd
			if fields[1] == "cd" {
				// if 'cd ..' go up one folder (parent)
				if fields[2] == ".." {
					current_folder = current_folder.parent
				} else {
					// check if folder exists in all_folders
					folder, ok := current_folder.subfolders[fields[2]]
					if ok {
						// if folder exists, change current_folder to it
						current_folder = folder
					} else {
						// if folder doesn't exist, create it, switch to it
						new_folder := newFolder(fields[2], current_folder)
						current_folder.subfolders[fields[2]] = new_folder
						current_folder = new_folder
					}

				}
			}
		} else { // line is a folder or file from ls command
			// ignore any lines that start with dir since they will be
			// entered as a subfolder
			if fields[0] != "dir" {
				// parse first number as int and add it to the current folder's size
				size, err := strconv.Atoi(fields[0])
				if err != nil {
					return "", err
				}
				current_folder.size += size
			}
		}
		//fmt.Printf("%s  -  %+v, %p\n", line, current_folder, current_folder)
	}
	total_size = 0
	folder_sizes = []int{}
	addSizes(root)
	folder_sizes = append(folder_sizes, root.size)

	total_space := 70000000
	required_unused_space := 30000000
	max_used_space := total_space - required_unused_space
	used_space := root.size
	min_deleted_space := used_space - max_used_space

	// sort folder_sizes and find first value over required_space
	sort.Slice(folder_sizes, func(i, j int) bool {
		return folder_sizes[i] < folder_sizes[j]
	})

	for _, size := range folder_sizes {
		if size > min_deleted_space {
			return fmt.Sprint(size), nil
		}
	}

	// no folder is large enough
	return "", fmt.Errorf("no folder is large enough")
}

func main() {
	// get input
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input_str := string(input)

	// part 1
	ans, err := part_1(input_str)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)

	// part 2
	ans, err = part_2(input_str)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans)
}
