from tree_node import TreeNode
from intro import traversal


def test_traversal():
    #      1
    #     / \
    #    2   3
    #   / \
    #  4   5
    root = TreeNode(
        1,
        TreeNode(2, TreeNode(4), TreeNode(5)),
        TreeNode(3),
    )

    assert traversal(root) == [1, 2, 3, 4, 5]


def test_empty_tree():
    assert traversal(None) == []
