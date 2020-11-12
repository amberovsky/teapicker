import React, { Component } from "react";
import axios from "axios";
import Button from '@material-ui/core/Button';
import HighlightOff from '@material-ui/icons/HighlightOff';
import AddCircle from "@material-ui/icons/AddCircle";
import TextField from "@material-ui/core/TextField";
import PersonIcon from '@material-ui/icons/Person';
import TableBody from "@material-ui/core/TableBody";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import Table from "@material-ui/core/Table";
import TableContainer from "@material-ui/core/TableContainer";

let endpoint = "http://localhost:8080";


/**
 * List of all employees
 */

export default class Employees extends Component {
    constructor(props) {
        super(props);

        this.handleChange = this.handleChange.bind(this);

        this.state = {
            newEmployee: "", // Employees' name when the "+" button is clicked
        }
    }

    handleChange(employees) {
        this.props.handleEmployeesChange(employees);
    }

    componentDidMount() {
        this.getEmployees();
    }

    getEmployees = () => {
        axios.get(endpoint + "/employee").then(res => {
            this.handleChange(res.data ? res.data : [])
        });
    };

    handleRemove = uuid => {
        axios.delete(endpoint + "/employee/" + uuid)
            .then(() => {
                this.getEmployees();
            });
    }

    handleNewEmployeeChange = event => {
        this.setState({
            newEmployee: event.target.value,
        })
    }

    handleAddSubmit = event => {
        event.preventDefault()
        axios
            .post(endpoint + "/employee", {name: this.state.newEmployee})
            .then(() => {
                this.getEmployees();
                this.setState({
                    newEmployee: "",
                })
            });
    }

    render() {
        return (
            <TableContainer>
                <Table size="small">
                    <TableBody>
                        {this.props.employees.map(employee => {
                            return (
                                <TableRow key={employee.uuid}>
                                    <TableCell>
                                        <div style={{display: 'flex', alignItems: 'center'}}>
                                            <PersonIcon />&nbsp;&nbsp;{employee.name}
                                        </div>
                                    </TableCell>
                                    <TableCell align={"right"}>
                                        <Button
                                            variant="text"
                                            color="secondary"
                                            startIcon={<HighlightOff />}
                                            onClick={() => {this.handleRemove(employee.uuid)}}
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
                                        label="New employee"
                                        variant="outlined"
                                        size="small"
                                        required={true}
                                        value={this.state.newEmployee}
                                        onChange={this.handleNewEmployeeChange}
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
