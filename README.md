# birthday-jamz

The goal is to generate a playlist for the members of a team

Input

1. the birth dates of the team members

e.g.:

1985-10-27
1996-12-23
1989-01-16

Process

Look for the top songs for on the birthday of each team member for every year since their birth

Output

1985-10-27
1986-10-27
1987-10-27
....
1996-10-27
1996-12-23
1997-10-27
1997-12-23





An API built with [Go Fiber](https://docs.gofiber.io/) to generate playlists from Billboard rankings for a given birthday.

## Running the app
1. Install Go 1.21.4 or higher
2. Clone repo
3. Start app locally with `go run .`

## Integration

We are using this API for our Billboard chart data: https://rapidapi.com/LDVIN/api/billboard-api
