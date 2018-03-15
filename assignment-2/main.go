package main

import (
	"fmt"
)

type node struct {
	question string
	result   string
	yes      *node
	no       *node
}

func main() {
	// Terminal nodes with no leaves but result with the decision
	passExamTerminalNode := &node{
		result: "You should be able to pass the exam",
	}
	easilyFailTerminalNode := &node{
		result: "You could easily fail the exam",
	}

	// Tree of nodes - made in the order they were represented in the assignment
	// some of the nodes er terminated faster than others because there is no reason
	// to go further when some criteria is furfilled
	tree := &node{
		question: "Read textbook",
		yes: &node{
			question: "Hand ins made in time",
			yes: &node{
				question: "Attend lectures",
				yes: &node{
					question: "Make exercises",
					yes:      passExamTerminalNode,
					no:       passExamTerminalNode, // answers4
				},
				no: &node{
					question: "Make exercises",
					yes:      passExamTerminalNode,
					no:       easilyFailTerminalNode, // answers2
				},
			},
			no: easilyFailTerminalNode, // answers1
		},
		no: &node{
			question: "Hand ins made in time",
			yes: &node{
				question: "Attend lectures",
				yes: &node{
					question: "Make exercises",
					yes:      passExamTerminalNode,
					no:       easilyFailTerminalNode,
				},
				no: easilyFailTerminalNode,
			},
			no: easilyFailTerminalNode, // answers3
		},
	}

	// slice of answers in the order they were represented in the assignment
	answers1 := []string{"yes", "no", "yes", "yes"}
	answers2 := []string{"yes", "yes", "no", "no"}
	answers3 := []string{"no", "no", "yes", "yes"}
	answers4 := []string{"yes", "yes", "yes", "no"}
	// decide returns the decision, and is given a tree(current node), the answers and the index of the current answer
	result1 := decide(tree, answers1, 0)
	result2 := decide(tree, answers2, 0)
	result3 := decide(tree, answers3, 0)
	result4 := decide(tree, answers4, 0)

	fmt.Println("result:", result1) // "You could easily fail the exam"
	fmt.Println("result:", result2) // "You could easily fail the exam"
	fmt.Println("result:", result3) // "You could easily fail the exam"
	fmt.Println("result:", result4) // "You should be able to pass the exam"
}

func decide(currentNode *node, answers []string, i int) string {
	// if there is a result, return the decision
	if currentNode.result != "" {
		return currentNode.result
	}

	fmt.Println("question", currentNode.question)
	fmt.Println("answer", answers[i])

	// if the answer is yes, return the "yes" node, else return the "no" node
	// and increment the answer index, after it's checked
	if answers[i] == "yes" {
		i++
		return decide(currentNode.yes, answers, i)
	}
	i++
	return decide(currentNode.no, answers, i)
}