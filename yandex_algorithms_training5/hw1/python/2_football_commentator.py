score1 = input().split(':')
score2 = input().split(':')

home_or_away_loc = int(input())

score1 = [int(goal) for goal in score1]
score2 = [int(goal) for goal in score2]


def count_goals(first: list, second: list, home_or_away: int) -> int:
    if home_or_away == 1:
        first_team_home = first[0]
        first_team_away = second[0]
        second_team_home = second[1]
        second_team_away = first[1]
    else:
        first_team_home = second[0]
        first_team_away = first[0]
        second_team_home = first[1]
        second_team_away = second[1]

    first_team_goals = first_team_home + first_team_away
    second_team_goals = second_team_home + second_team_away

    if first_team_goals > second_team_goals:
        return 0

    if first_team_goals == second_team_goals == 0:
        return 1

    equalize_goals = second_team_goals - first_team_goals

    if home_or_away == 1:
        if equalize_goals + first_team_away > second_team_away:
            return equalize_goals
        else:
            return equalize_goals + 1
    else:
        if first_team_away > second_team_away:
            return equalize_goals
        else:
            return equalize_goals + 1


print(count_goals(score1, score2, home_or_away_loc))
