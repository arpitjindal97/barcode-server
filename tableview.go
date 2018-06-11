package main

import (
	"fmt"
	"github.com/therecipe/qt/core"
)

func init() { CustomTableModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomTableModel") }

const (
	ScanResults = int(core.Qt__UserRole) + 1<<iota
)

type TableItem struct {
	ScanResult string
}

type CustomTableModel struct {
	core.QAbstractTableModel

	_ func() `constructor:"init"`

	_ func()                          `signal:"remove,auto"`
	_ func(item TableItem)            `signal:"add,auto"`
	_ func(index int, item TableItem) `signal:"edit,auto"`
	_ func()                          `signal:"start,auto"`

	modelData []TableItem
}

func (m *CustomTableModel) init() {
	m.ConnectRoleNames(m.roleNames)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectData(m.data)
}

func (m *CustomTableModel) roleNames() map[int]*core.QByteArray {
	return map[int]*core.QByteArray{
		ScanResults: core.NewQByteArray2("Scan Results", -1),
	}
}

func (m *CustomTableModel) rowCount(*core.QModelIndex) int {
	return len(m.modelData)
}

func (m *CustomTableModel) columnCount(*core.QModelIndex) int {
	return 1
}

func (m *CustomTableModel) data(index *core.QModelIndex, role int) *core.QVariant {
	item := m.modelData[index.Row()]
	switch role {
	case ScanResults:
		return core.NewQVariant14(item.ScanResult)
	}
	return core.NewQVariant()
}

func (m *CustomTableModel) remove() {
	if len(m.modelData) == 0 {
		return
	}
	m.BeginRemoveRows(core.NewQModelIndex(), len(m.modelData)-1, len(m.modelData)-1)
	m.modelData = m.modelData[:len(m.modelData)-1]
	m.EndRemoveRows()
}

func (m *CustomTableModel) add(item TableItem) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.modelData), len(m.modelData))
	m.modelData = append(m.modelData, item)
	m.EndInsertRows()
}

func (m *CustomTableModel) edit(index int, item TableItem) {
	if len(m.modelData) == 0 {
		return
	}
	m.modelData[index] = item
	m.DataChanged(m.Index(index, 0, core.NewQModelIndex()), m.Index(index, 1, core.NewQModelIndex()),
		[]int{ScanResults})
}

func (m *CustomTableModel) start() {
	//go StartProcess()
	fmt.Println("start button clicked")
}
