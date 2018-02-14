package ast // import "ttp.sh/flatpack/pkg/ast"

import "ttp.sh/flatpack/pkg/token"

// ----------------------------------------------------------------------------
// Interfaces
//
// There are 3 main classes of nodes: Expressions and type nodes,
// statement nodes, and declaration nodes. The node names usually
// match the corresponding Go spec production names to which they
// correspond. The node fields correspond to the individual parts
// of the respective productions.
//
// All nodes contain position information marking the beginning of
// the corresponding source text segment; it is accessible via the
// Pos accessor method. Nodes may contain additional position info
// for language constructs where comments may be found between parts
// of the construct (typically any larger, parenthesized subpart).
// That position information is needed to properly position comments
// when printing the construct.

// Node All node types implement the Node interface.
type Node interface {
	Pos() token.Pos
	End() token.Pos
}

// Expr All expression nodes implement the Expr interface.
type Expr interface {
	Node
	exprNode()
}

// Stmt All statement nodes implement the Stmt interface.
type Stmt interface {
	Node
	stmtNode()
}

// Decl All declaration nodes implement the Decl interface.
type Decl interface {
	Node
	declNode()
}

type (
	// A BasicLit node represents a literal of basic type.
	BasicLit struct {
		ValuePos token.Pos   // literal position
		Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
		Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
	}

	// An ExprStmt node represents a (stand-alone) expression
	// in a statement list.
	//
	ExprStmt struct {
		X Expr // expression
	}
)

// Pos and End implementations for statement nodes.

func (s *ExprStmt) Pos() token.Pos { return s.X.Pos() }
func (s *ExprStmt) End() token.Pos { return s.X.End() }

func (*ExprStmt) stmtNode() {}
