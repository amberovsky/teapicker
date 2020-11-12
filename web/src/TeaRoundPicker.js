import React, { Component } from "react";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import Avatar from "@material-ui/core/Avatar";
import Typography from "@material-ui/core/Typography";
import Table from "@material-ui/core/Table";


/**
 * Render the picked team member
 */

export default class TeaRoundPicker extends Component {
    render() {
        const teaMakerUuid = this.props.teaMaker
        const employees = this.props.employees

        let avatar = <div></div>
        let teaMaker = <div></div>

        if (teaMakerUuid !== "-1") {
            avatar = <Avatar src={"https://thispersondoesnotexist.com/image?q=" + teaMakerUuid} style={{width: '100px', height: '100px', marginTop: '10px'}} />

            const employeesLength = employees.length
            for (let i = 0; i < employeesLength; i++) {
                if (employees[i].uuid === teaMakerUuid) {
                    teaMaker =
                        <Typography variant="h5" align="center" color="textSecondary" component="p">
                            It's <b>{employees[i].name}</b>
                        </Typography>

                    break;
                }
            }
        }

        return (
            <Table>
                <TableBody>
                    <TableRow>
                        <TableCell style={{borderBottom: "0px", width: "100px"}}>
                            {avatar}
                        </TableCell>
                        <TableCell style={{borderBottom: "0px"}} align={"center"}>
                            {teaMaker}
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        );
    }
}
