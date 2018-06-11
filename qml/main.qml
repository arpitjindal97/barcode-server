import QtQuick 2.7
import QtQuick.Controls 1.4
import QtQuick.Controls 2.1
import CustomQmlTypes 1.0

Item {
    id:window
    width: 400
    height:400

    Rectangle {
        id: rectangle
        x: 0
        y: 0
        width: 400
        height: 400
        color: "#ffffff"

        Rectangle {
            id: rectangle1
            x: 0
            y: 0
            width: 400
            height: 60
            color: "#0f35f4"

            Text {
                id: text2
                x: 34
                y: 12
                width: 332
                height: 36
                color: "#ffffff"
                text: qsTr("Barcode Scan Reciever")
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                font.underline: false
                font.pixelSize: 22
                font.weight: Font.Light
                font.bold: true
                font.family: "Arial"
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
            }
        }

        Rectangle {
            id: rectangle2
            x: 0
            y: 60
            width: 208
            height: 266

            TableView {
                id: tableView
                x: 8
                y: 8
                width: 200
                height: 258
                horizontalScrollBarPolicy: 2
                // @disable-check M300
                model: CustomTableModel{}

                TableViewColumn {
                        role: "results"
                        title: "Scan Results"
                    }
            }
        }

        Rectangle {
            id: rectangle3
            x: 208
            y: 60
            width: 192
            height: 266
            color: "#ffffff"
        }

        Rectangle {
            id: rectangle4
            x: 0
            y: 327
            width: 400
            height: 52
            color: "#ffffff"

            Button {
                id: button
                x: 51
                y: 13
                width: 101
                height: 28
                text: qsTr("Start")
                onClicked: tableView.model.start()
                anchors.verticalCenter: parent.verticalCenter
            }

            Button {
                id: button1
                x: 255
                y: 13
                width: 99
                height: 28
                text: qsTr("Update")
                anchors.verticalCenter: parent.verticalCenter
            }
        }

        Text {
            id: text1
            x: 242
            y: 385
            width: 158
            height: 15
            text: qsTr("Created By Arpit Agarwal")
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 3
            font.pixelSize: 12
        }

    }
}
