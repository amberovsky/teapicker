import React, { Component } from "react";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import axios from "axios";

let endpoint = "http://localhost:8080";


/**
 * Which team is currently selected in the "Team  members" section
 */

export default class MembersSelect extends Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);
    }

    handleChange(event) {
        this.props.handleSelectedTeamForMembersChange(event.target.value)

        if (event.target.value !== "-1") {
            axios.get(endpoint + "/member?team_uuid=" + event.target.value).then(res => {
                this.props.handleTeamMembersChange(res.data ? res.data : []);
            });
        }
    }

    render() {
        const teams = this.props.teams
        const selectedTeamForMembers = this.props.selectedTeamForMembers

        return (
            <FormControl fullWidth={true}>
                <InputLabel id="team-selector-label">Select team</InputLabel>
                <Select
                    labelId="team-selector-label"
                    id="team-selector"
                    value={selectedTeamForMembers}
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
        );
    }
}
