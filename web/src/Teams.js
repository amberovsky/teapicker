import React, { Component } from "react";
import axios from "axios";
import TableContainer from "@material-ui/core/TableContainer";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import TextField from "@material-ui/core/TextField";
import Button from "@material-ui/core/Button";
import AddCircle from "@material-ui/icons/AddCircle";
import GroupIcon from "@material-ui/icons/Group";
import HighlightOff from "@material-ui/icons/HighlightOff";

let endpoint = "http://localhost:8080";


/**
 * List of all teams
 */

export default class Teams extends Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);

        this.state = {
            newTeam: "",
        }
    }

    handleChange(teams) {
        this.props.handleTeamsChange(teams);
    }

    componentDidMount() {
        this.getTeams();
    }

    getTeams = () => {
        axios.get(endpoint + "/team").then(res => {
            this.handleChange(res.data ? res.data : [])
        });
    };

    handleRemove = uuid => {
        axios.delete(endpoint + "/team/" + uuid)
            .then(() => {
                this.getTeams();
            });
    }

    handleNewTeamChange = event => {
        this.setState({
            newTeam: event.target.value,
        })
    }

    handleAddSubmit = event => {
        event.preventDefault()
        axios
            .post(endpoint + "/team", {name: this.state.newTeam})
            .then(() => {
                this.getTeams();
                this.setState({
                    newTeam: "",
                })
            });
    }

    render() {
        const teams = this.props.teams

        return (
            <TableContainer>
                <Table size="small">
                    <TableBody>
                        {teams.map(team => {
                            return (
                                <TableRow key={team.uuid}>
                                    <TableCell>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                                            <GroupIcon />&nbsp;&nbsp;{team.name}
                                        </div>
                                    </TableCell>
                                    <TableCell align={"right"}>
                                        <Button
                                            variant="text"
                                            color="secondary"
                                            startIcon={<HighlightOff />}
                                            onClick={() => {this.handleRemove(team.uuid)}}
                                        />
                                    </TableCell>
                                </TableRow>
                            );
                        })}
                    </TableBody>
                </Table>
                <form autoComplete="off" onSubmit={this.handleAddSubmit} action="#">
                    <Table size="small">
                        <TableBody>
                            <TableRow key="new">
                                <TableCell>
                                    <TextField
                                        label="New team"
                                        variant="outlined"
                                        size="small"
                                        required={true}
                                        value={this.state.newTeam}
                                        onChange={this.handleNewTeamChange}
                                    />
                                </TableCell>
                                <TableCell align={"right"}>
                                    <Button
                                        type="submit"
                                        variant="text"
                                        color="primary"
                                        startIcon={<AddCircle />}
                                    />
                                </TableCell>
                            </TableRow>

                        </TableBody>
                    </Table>
                </form>
            </TableContainer>
        );
    }
}
