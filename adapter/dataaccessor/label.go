package dataaccessor

import "github.com/pepnova-9/cqrs-sample/domain"

type Label struct {
}

// Create はラベルを作成する
func (l *Label) Create(entity domain.Label) error {
	_ = `
	INSERT xxx INTO labels;
`
	return nil
}

// Find はラベルを取得する
func (l *Label) Find(id domain.LabelID) (domain.Label, error) {
	_ = `
	SELECT * FROM labels WHERE id = ${id};
`
	// DataModelからEntityに変換する

	return domain.Label{}, nil
}
