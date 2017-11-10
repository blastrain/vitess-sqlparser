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

func convertFromTruncateTableStmt(stmt *ast.TruncateTableStmt) Statement {
	return &TruncateTable{Table: TableName{Name: TableIdent{v: stmt.Table.Name.String()}}}
}

func convertTiDBStmtToVitessOtherAdmin(stmts []ast.StmtNode, admin *OtherAdmin) Statement {
	for _, stmt := range stmts {
		switch adminStmt := stmt.(type) {
		case *ast.TruncateTableStmt:
			return convertFromTruncateTableStmt(adminStmt)
		default:
			return admin
		}
	}
	return nil
}

func convertTiDBStmtToVitessDDL(stmts []ast.StmtNode, ddl *DDL) Statement {
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
