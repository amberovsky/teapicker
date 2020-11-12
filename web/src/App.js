import React from "react";
import Container from "@material-ui/core/Container";
import Typography from '@material-ui/core/Typography';
import Toolbar from '@material-ui/core/Toolbar';
import OutdoorGrill from '@material-ui/icons/OutdoorGrill';
import TeaRoundPicker from "./TeaRoundPicker";
import Employees from "./Employees";
import Teams from "./Teams";
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import Link from "@material-ui/core/Link";
import MembersSelect from "./MembersSelect";
import {withStyles} from "@material-ui/core/styles";
import Members from "./Members";
import TeamSelect from "./TeamSelect";


/**
 * Main entrypoint
 */

const styles = theme => ({
    '@global': {
        ul: {
            margin: 0,
            padding: 0,
            listStyle: 'none',
        },
    },
    appBar: {
        borderBottom: `1px solid ${theme.palette.divider}`,
    },
    toolbar: {
        flexWrap: 'wrap',
    },
    heroContent: {
        padding: theme.spacing(8, 0, 10),
    },
    employeesCardHeader: {
        backgroundColor:
            theme.palette.info.light
    },
    teamsCardHeader: {
        backgroundColor:
            theme.palette.secondary.light
    },
    membersCardHeader: {
        backgroundColor:
            theme.palette.success.light
    },
    footer: {
        borderTop: `1px solid ${theme.palette.divider}`,
        marginTop: theme.spacing(8),
        paddingTop: theme.spacing(3),
        paddingBottom: theme.spacing(3),
        [theme.breakpoints.up('sm')]: {
            paddingTop: theme.spacing(6),
            paddingBottom: theme.spacing(6),
        },
    },
});

function Copyright() {
    return (
        <Typography variant="body2" color="textSecondary" align="center">
            {'Copyright © '}
            <Link color="inherit" href="#">
                Anton Zagorskii
            </Link>{' '}
            {new Date().getFullYear()}
        </Typography>
    );
}

class App extends React.Component {

    constructor(props) {
        super(props);

        this.handleTeamsChange = this.handleTeamsChange.bind(this);
        this.handleEmployeesChange = this.handleEmployeesChange.bind(this);
        this.handleSelectedTeamForMembersChange = this.handleSelectedTeamForMembersChange.bind(this);
        this.handleSelectedTeamChange = this.handleSelectedTeamChange.bind(this);
        this.handleTeamMembersChange = this.handleTeamMembersChange.bind(this);
        this.handleTeaMakerChange = this.handleTeaMakerChange.bind(this);

        this.state = {
            teams: [], // List of all teams
            employees: [], // List of all employees
            selectedTeamForMembers: "-1", // Which team is currently selected in the "Team  members" section
            teamMembers: [], // List of team members for the selected team
            selectedTeam: "-1", // Which team is selected for the next tea round
            teaMaker: "-1", // Picked employee to do the next tea round
        };
    }


    /**
     * Handlers for user interaction
     */

    handleTeamsChange(teams) {
        let foundSelectedTeamForMembers = teams.some(team => {
            return (team.uuid === this.state.selectedTeamForMembers)
        })

        let foundSelectedTeam = teams.some(team => {
            return (team.uuid === this.state.selectedTeam)
        })

        this.setState({
            teams: teams,
            selectedTeam: foundSelectedTeam ? this.state.selectedTeam : "-1",
            teaMaker: foundSelectedTeam ? this.state.teaMaker : "-1",
            selectedTeamForMembers: foundSelectedTeamForMembers ? this.state.selectedTeamForMembers : "-1",
        });
    }

    handleEmployeesChange(employees) {
        let teaMakerNotAffected = employees.some(employee => {
            return (employee.uuid === this.state.teaMaker)
        })

        this.setState({
            employees: employees,
            teaMaker: teaMakerNotAffected ? this.state.teaMaker : "-1",
        });
    }

    handleSelectedTeamForMembersChange(selectedTeamForMembers) {
        this.setState({selectedTeamForMembers: selectedTeamForMembers});
    }

    handleSelectedTeamChange(selectedTeam) {
        this.setState({
            teaMaker: selectedTeam !== this.state.selectedTeam  ? "-1" : this.state.teaMaker,
            selectedTeam: selectedTeam
        });
    }

    handleTeamMembersChange(teamMembers) {
        let teaMakerNotAffected = teamMembers.some(teamMember => {
            return (teamMember.employee_uuid === this.state.teaMaker)
        })

        this.setState({
            teamMembers: teamMembers,
            teaMaker: teaMakerNotAffected ? this.state.teaMaker : "-1"
        });
    }

    handleTeaMakerChange(teaMaker) {
        this.setState({teaMaker: teaMaker});
    }

    render() {
        const { classes } = this.props;

        const teams = this.state.teams
        const employees = this.state.employees
        const selectedTeamForMembers = this.state.selectedTeamForMembers
        const selectedTeam = this.state.selectedTeam
        const teamMembers = this.state.teamMembers
        const teaMaker = this.state.teaMaker

        return (
            <div>
                <CssBaseline />
                <AppBar position="relative" className={classes.appBar}>
                    <Toolbar className={classes.toolbar}>
                        <OutdoorGrill />
                        <Typography variant="h6" color="inherit" noWrap>
                            &nbsp;Lumio Research & Development Centre
                        </Typography>
                    </Toolbar>
                </AppBar>

                <Container maxWidth="sm" className={classes.heroContent}>
                    <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                        The Tea Round Picker
                    </Typography>

                    <Card>
                        <CardContent>
                            <TeamSelect
                                teams={teams}
                                selectedTeam={selectedTeam}
                                handleSelectedTeamChange={this.handleSelectedTeamChange}
                                handleTeaMakerChange={this.handleTeaMakerChange}
                            />

                            <TeaRoundPicker
                                employees={employees}
                                teaMaker={teaMaker}
                            />

                        </CardContent>
                    </Card>
                </Container>

                <main>
                    <Container maxWidth="lg" component="main">
                        <Grid container spacing={5} alignItems="flex-start">
                            <Grid item key="Employees" xs={12} sm={12} md={4}>
                                <Card>
                                    <CardHeader
                                        title="Employees"
                                        subheader="all tea drinkers at Lumio"
                                        titleTypographyProps={{align: 'center'}}
                                        subheaderTypographyProps={{align: 'center'}}
                                        className={classes.employeesCardHeader}
                                    />
                                    <CardContent>
                                        <Employees
                                            employees={employees}
                                            handleEmployeesChange={this.handleEmployeesChange}
                                        />
                                    </CardContent>
                                </Card>
                            </Grid>

                            <Grid item key="Teams" xs={12} sm={12} md={4}>
                                <Card>
                                    <CardHeader
                                        title="Teams"
                                        subheader="any possible team"
                                        titleTypographyProps={{align: 'center'}}
                                        subheaderTypographyProps={{align: 'center'}}
                                        className={classes.teamsCardHeader}
                                    />
                                    <CardContent>
                                        <Teams teams={teams} handleTeamsChange={this.handleTeamsChange} />
                                    </CardContent>
                                </Card>
                            </Grid>

                            <Grid item key="MembersSelect" xs={12} sm={12} md={4}>
                                <Card>
                                    <CardHeader
                                        title="Team members"
                                        subheader="any employee can join any teams"
                                        titleTypographyProps={{align: 'center'}}
                                        subheaderTypographyProps={{align: 'center'}}
                                        className={classes.membersCardHeader}
                                    />
                                    <CardContent>
                                        <MembersSelect
                                            teams={teams}
                                            selectedTeamForMembers={selectedTeamForMembers}
                                            handleSelectedTeamForMembersChange={this.handleSelectedTeamForMembersChange}
                                            handleTeamMembersChange={this.handleTeamMembersChange}
                                        />

                                        <Members
                                            selectedTeamForMembers={selectedTeamForMembers}
                                            employees={employees}
                                            teamMembers={teamMembers}
                                            handleTeamMembersChange={this.handleTeamMembersChange}
                                        />
                                    </CardContent>
                                </Card>
                            </Grid>

                        </Grid>
                    </Container>
                    <Copyright/>
                </main>
            </div>
        );
    }
}

export default withStyles(styles)(App);
