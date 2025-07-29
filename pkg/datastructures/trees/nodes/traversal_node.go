package nodes

import (
	"cmp"
	"fmt"

	"ds-algorithms/pkg/datastructures/array"
)

type TraversableNode[E cmp.Ordered] interface {
	GetLeft() TraversableNode[E]
	GetRight() TraversableNode[E]
	GetValue() E
	IsNil() bool
}

func PreOrderTraversal[E cmp.Ordered](node TraversableNode[E], resultArr *array.ArrayList[E]) error {
	if node.IsNil() {
		return nil
	}

	err := resultArr.Add(resultArr.Size(), node.GetValue())
	if err != nil {
		return fmt.Errorf("error adding element %v to result list: %w", node.GetValue(), err)
	}

	err = PreOrderTraversal(node.GetLeft(), resultArr)
	if err != nil {
		return err
	}

	err = PreOrderTraversal(node.GetRight(), resultArr)
	if err != nil {
		return err
	}

	return nil
}

func PostOrderTraversal[E cmp.Ordered](node TraversableNode[E], resultArr *array.ArrayList[E]) error {
	if node.IsNil() {
		return nil
	}
	err := PostOrderTraversal(node.GetLeft(), resultArr)
	if err != nil {
		return err
	}

	err = PostOrderTraversal(node.GetRight(), resultArr)
	if err != nil {
		return err
	}

	err = resultArr.Add(resultArr.Size(), node.GetValue())
	if err != nil {
		return fmt.Errorf("error adding element %v to result list: %w", node.GetValue(), err)
	}

	return nil
}

func InOrderTraversal[E cmp.Ordered](node TraversableNode[E], resultArr *array.ArrayList[E]) error {
	if node.IsNil() {
		return nil
	}
	err := InOrderTraversal(node.GetLeft(), resultArr)
	if err != nil {
		return err
	}

	err = resultArr.Add(resultArr.Size(), node.GetValue())
	if err != nil {
		return fmt.Errorf("error adding element %v to result list: %w", node.GetValue(), err)
	}

	err = InOrderTraversal(node.GetRight(), resultArr)
	if err != nil {
		return err
	}

	return nil
}
