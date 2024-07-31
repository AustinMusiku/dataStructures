package B_tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {

	t.Run("empty tree", func(t *testing.T) {
		t.Parallel()
		tree, _ := New(4)
		found, ok := tree.Search([]byte("a"))

		assert.False(t, ok, "key should not be found")
		assert.Nil(t, found, "value should be nil")
	})

	t.Run("multiple keys", func(t *testing.T) {
		t.Parallel()

		tree, _ := New(4)
		keys := []string{"a", "j", "m", "z", "t"}

		for _, key := range keys {
			tree.Insert([]byte(key))
		}

		found1, ok1 := tree.Search([]byte("a"))
		found2, ok2 := tree.Search([]byte("j"))
		found3, ok3 := tree.Search([]byte("m"))
		found4, ok4 := tree.Search([]byte("z"))
		found5, ok5 := tree.Search([]byte("t"))

		found6, ok6 := tree.Search([]byte("b"))
		found7, ok7 := tree.Search([]byte("k"))

		assert.True(t, ok1, "key should be found")
		assert.Equal(t, []byte("a"), found1, "value should be 'a'")

		assert.True(t, ok2, "key should be found")
		assert.Equal(t, []byte("j"), found2, "value should be 'j'")

		assert.True(t, ok3, "key should be found")
		assert.Equal(t, []byte("m"), found3, "value should be 'm'")

		assert.True(t, ok4, "key should be found")
		assert.Equal(t, []byte("z"), found4, "value should be 'z'")

		assert.True(t, ok5, "key should be found")
		assert.Equal(t, []byte("t"), found5, "value should be 't'")

		assert.False(t, ok6, "key should not be found")
		assert.Nil(t, found6, "value should be nil")

		assert.False(t, ok7, "key should not be found")
		assert.Nil(t, found7, "value should be nil")
	})
}

func Test_Insert(t *testing.T) {

	t.Run("multiple keys", func(t *testing.T) {
		t.Parallel()

		tree, _ := New(4)
		keys := []string{"10", "20", "40", "50", "60", "70", "80", "30", "35", "05", "15"}
		for _, key := range keys {
			tree.Insert([]byte(key))
		}

		// root
		assert.Equal(t, 1, len(tree.root.inodes), "root should have 1 inodes")
		assert.Equal(t, 2, len(tree.root.children), "root should have 2 children")

		// children
		assert.Equal(t, 2, len(tree.root.children[0].inodes), "child 1 should have 2 inodes")
		assert.Equal(t, 3, len(tree.root.children[0].children), "child 1 should have 3 children")

		assert.Equal(t, 1, len(tree.root.children[1].inodes), "child 2 should have 1 inodes")
		assert.Equal(t, 2, len(tree.root.children[1].children), "child 2 should have 2 children")

		// grandchildren
		assert.Equal(t, 2, len(tree.root.children[0].children[0].inodes), "grandchild 1 should have 2 inodes")
		assert.Equal(t, 0, len(tree.root.children[0].children[0].children), "grandchild 1 should have 0 children")

		assert.Equal(t, 1, len(tree.root.children[0].children[1].inodes), "grandchild 2 should have 1 inodes")
		assert.Equal(t, 0, len(tree.root.children[0].children[1].children), "grandchild 2 should have 0 children")

		assert.Equal(t, 1, len(tree.root.children[0].children[2].inodes), "grandchild 3 should have 1 inodes")
		assert.Equal(t, 0, len(tree.root.children[0].children[2].children), "grandchild 3 should have 0 children")

		assert.Equal(t, 2, len(tree.root.children[1].children[0].inodes), "grandchild 4 should have 2 inodes")
		assert.Equal(t, 0, len(tree.root.children[1].children[0].children), "grandchild 4 should have 0 children")

		assert.Equal(t, 1, len(tree.root.children[1].children[1].inodes), "grandchild 5 should have 1 inodes")
		assert.Equal(t, 0, len(tree.root.children[1].children[1].children), "grandchild 5 should have 0 children")
	})

}

func Test_Delete(t *testing.T) {
	t.Run("multiple key removals", func(t *testing.T) {
		t.Parallel()

		tree, _ := New(4)
		keys := []string{"10", "20", "40", "50", "60", "70", "80", "30", "35", "05", "15"}
		for _, key := range keys {
			tree.Insert([]byte(key))
		}

		// Remove keys and check structure
		keysToRemove := []string{"70", "35", "15"}
		for _, key := range keysToRemove {
			err := tree.Delete([]byte(key))
			assert.NoError(t, err, "Delete should not return an error")
		}

		// Verify tree structure after removals

		// root
		assert.Equal(t, 1, len(tree.root.inodes), "root should have 1 inode")
		assert.Equal(t, string(tree.root.inodes[0]), "40", "root should have 40 as the only inode key")
		assert.Equal(t, 2, len(tree.root.children), "root should have 2 children")

		// children
		assert.Equal(t, 1, len(tree.root.children[0].inodes), "child 1 should have 1 inodes")
		assert.Equal(t, 2, len(tree.root.children[0].children), "child 1 should have 1 children")

		assert.Equal(t, 1, len(tree.root.children[1].inodes), "child 2 should have 1 inode")
		assert.Equal(t, 2, len(tree.root.children[1].children), "child 2 should have 2 children")

		// grandchildren
		assert.Equal(t, 1, len(tree.root.children[0].children[0].inodes), "grandchild 1 should have 1 inodes")
		assert.Equal(t, 0, len(tree.root.children[0].children[0].children), "grandchild 1 should have 0 children")

		assert.Equal(t, 2, len(tree.root.children[0].children[1].inodes), "grandchild 2 should have 2 inode")
		assert.Equal(t, 0, len(tree.root.children[0].children[1].children), "grandchild 2 should have 0 children")

		assert.Equal(t, 1, len(tree.root.children[1].children[0].inodes), "grandchild 3 should have 1 inode")
		assert.Equal(t, 0, len(tree.root.children[1].children[0].children), "grandchild 3 should have 0 children")

		assert.Equal(t, 1, len(tree.root.children[1].children[1].inodes), "grandchild 4 should have 1 inode")
		assert.Equal(t, 0, len(tree.root.children[1].children[1].children), "grandchild 4 should have 0 children")

		// Verify specific key removals
		for _, key := range keysToRemove {
			found, ok := tree.Search([]byte(key))
			assert.False(t, ok, fmt.Sprintf("Key %s should not be found after removal", key))
			assert.Nil(t, found, fmt.Sprintf("Value for key %s should be nil after removal", key))
		}

		// Verify remaining keys
		remainingKeys := []string{"05", "10", "20", "30", "40", "50", "60", "80"}
		for _, key := range remainingKeys {
			found, ok := tree.Search([]byte(key))
			assert.True(t, ok, fmt.Sprintf("Key %s should still be in the tree", key))
			assert.Equal(t, []byte(key), found, fmt.Sprintf("Value for key %s should still be %s", key, key))
		}
	})
}
