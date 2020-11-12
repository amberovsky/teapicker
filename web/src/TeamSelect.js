import React, { Component } from "react";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import {Table} from "@material-ui/core";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import Button from "@material-ui/core/Button";
import RepeatIcon from "@material-ui/icons/Repeat";
import axios from "axios";

let endpoint = "http://localhost:8080";


/**
 * Which team is selected for the next tea round
 */

export default class TeamSelect extends Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(event) {
        this.props.handleSelectedTeamChange(event.target.value)
    }

    handlePickClick = () => {
        axios.post(endpoint + "/pick/" + this.props.selectedTeam).then(res => {
            if (res.data) {
                axios.get(endpoint + "/member/" + res.data.member_uuid).then(res => {
                    this.props.handleTeaMakerChange(res.data.employee_uuid)
                });
            }
        });
    }

    render() {
        const teams = this.props.teams
        const selectedTeam = this.props.selectedTeam

        return (
            <Table>
                <TableBody>
                    <TableRow>
                        <TableCell>
                            <FormControl fullWidth={true}>
                                <InputLabel>Select team</InputLabel>
                                <Select
                                    labelId="team-selector-label"
                                    id="team-selector"
                                    value={selectedTeam}
                                    onChange={this.handleChange}
                                >
                                    <MenuItem value="-1">
                                        <em>None</em>
                                    </MenuItem>
                                    {teams.map(team => {
                                        return (
                                            <MenuItem value={team.uuid} key={team.uuid}>{team.name}</MenuItem>
                                        )
                                    })}
                                </Select>
                            </FormControl>
                        </TableCell>
                        <TableCell align={"right"} style={{width: "220px"}}>
                            <Button
                                variant="contained"
                                color="primary"
                                startIcon={<RepeatIcon />}
                                disabled={selectedTeam === "-1"}
                                onClick={this.handlePickClick}
                            >
                                Pick tea maker
                            </Button>
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        );
    }
}
