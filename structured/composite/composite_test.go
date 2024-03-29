package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := &file{name:"file1"}
	file2 := &file{name:"file2"}
	file3 := &file{name:"file3"}

	folder1 := &folder{name:"Folder1"}
	folder1.add(file1)

	folder2 := &folder{name:"Folder1"}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}
