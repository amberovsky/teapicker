import React, { Component } from "react";
import TableContainer from "@material-ui/core/TableContainer";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import GroupIcon from "@material-ui/icons/Group";
import Button from "@material-ui/core/Button";
import {Add, RemoveCircle} from "@material-ui/icons";
import axios from "axios";

let endpoint = "http://localhost:8080";


/**
 * List of team members for the selected team
 */

export default class Members extends Component {
    getTeamMembers = teamUuid => {
        axios.get(endpoint + "/member?team_uuid=" + teamUuid).then(res => {
            this.props.handleTeamMembersChange(res.data ? res.data : []);
        });
    }

    handleAdd = employeeUuid => {
        axios
            .post(endpoint + "/member", {
                team_uuid: this.props.selectedTeamForMembers,
                employee_uuid: employeeUuid,
            })
            .then(() => {
                this.getTeamMembers(this.props.selectedTeamForMembers);
            });
    }

    handleRemove = uuid => {
        axios.delete(endpoint + "/member/" + uuid)
            .then(() => {
                this.getTeamMembers(this.props.selectedTeamForMembers);
            });
    }

    render() {
        const selectedTeamForMembers = this.props.selectedTeamForMembers
        const teamMembers = this.props.teamMembers
        const employees = this.props.employees

        var inTeam = {}
        var notInTeam = []

        var membersLength = teamMembers.length

        if (selectedTeamForMembers !== "-1") {
            employees.forEach(employee => {
                let found = false;
                for (let i = 0; i < membersLength; i++) {
                    if (teamMembers[i].employee_uuid === employee.uuid) {
                        inTeam[teamMembers[i].uuid] = employee
                        found = true;
                        break;
                    }
                }

                if (!found) {
                    notInTeam.push(employee)
                }
            })
        }

        return (
            <TableContainer>
                <Table size="small">
                    <TableBody>
                        {Object.keys(inTeam).map(uuid => {
                            return (
                                <TableRow key={uuid}>
                                    <TableCell>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                                            <GroupIcon />&nbsp;&nbsp;{inTeam[uuid].name}
                                        </div>
                                    </TableCell>
                                    <TableCell align={"right"}>
                                        <Button
                                            variant="text"
                                            color="secondary"
                                            startIcon={<RemoveCircle />}
                                            onClick={() => {this.handleRemove(uuid)}}
                                        />
                                    </TableCell>
                                </TableRow>
                            );
                        })}
                    </TableBody>
                </Table>
                <Table size="small">
                    <TableBody>
                        {notInTeam.map(employee => {
                            return (
                                <TableRow key={employee.uuid}>
                                    <TableCell>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                                            <GroupIcon />&nbsp;&nbsp;{employee.name}
                                        </div>
                                    </TableCell>
                                    <TableCell align={"right"}>
                                        <Button
                                            variant="text"
                                            color="primary"
                                            startIcon={<Add />}
                                            onClick={() => {this.handleAdd(employee.uuid)}}
                                        />
                                    </TableCell>
                                </TableRow>
                            );
                        })}
                    </TableBody>
                </Table>
            </TableContainer>
        );
    }
}
