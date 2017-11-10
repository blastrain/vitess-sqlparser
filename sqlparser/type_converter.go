package sqlparser

import (
	"github.com/knocknote/vitess-sqlparser/tidbparser/ast"
)

func convertFromCreateTableStmt(stmt *ast.CreateTableStmt, ddl *DDL) Statement {
	var columns []*ColumnDef
	for _, col := range stmt.Cols {
		columns = append(columns, &ColumnDef{
			Name:  col.Name.Name.String(),
			Type:  col.Tp.String(),
			Elems: col.Tp.Elems,
		})
	}
	return &CreateTable{
		DDL:     ddl,
		Columns: columns,
	}
}

func convertTiDBStmtToVitessStmt(stmts []ast.StmtNode, ddl *DDL) Statement {
	for _, stmt := range stmts {
		switch ddlStmt := stmt.(type) {
		case *ast.CreateTableStmt:
			return convertFromCreateTableStmt(ddlStmt, ddl)
		default:
			return ddl
		}
	}
	return nil
}
