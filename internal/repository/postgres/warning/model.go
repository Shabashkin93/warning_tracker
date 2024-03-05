package warning

import (
	"database/sql"

	"github.com/Shabashkin93/warning_tracker/internal/domain/warning"
)

type pgWarning struct {
	ID        sql.NullString `db:"id"`
	Branch    sql.NullString `db:"branch"`
	Commit    sql.NullString `db:"commit"`
	Count     sql.NullInt32  `db:"count"`
	CreatedBy sql.NullString `db:"created_by"`
	CreatedAt sql.NullString `db:"created_at"`
}

func domainToPgWarning(in *warning.WarningCreate, out *pgWarning) {
	out.ID = sql.NullString{String: in.Id, Valid: true}
	out.Branch = sql.NullString{String: in.Branch, Valid: true}
	out.Commit = sql.NullString{String: in.Commit, Valid: true}
	out.Count = sql.NullInt32{Int32: int32(in.Count), Valid: true}
	out.CreatedBy = sql.NullString{String: in.CreatedBy, Valid: true}
	out.CreatedAt = sql.NullString{String: in.CreatedAt, Valid: true}
}

func warningToPg(in *pgWarning, out *warning.WarningCreate) {
	if in.ID.Valid {
		out.Id = in.ID.String
	}

	if in.Branch.Valid {
		out.Branch = in.Branch.String
	}

	if in.Commit.Valid {
		out.Commit = in.Commit.String
	}

	if in.Count.Valid {
		out.Count = int(in.Count.Int32)
	}

	if in.CreatedBy.Valid {
		out.CreatedBy = in.CreatedBy.String
	}

	if in.CreatedAt.Valid {
		out.CreatedAt = in.CreatedAt.String
	}
}
