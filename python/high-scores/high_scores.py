def latest(scores):
    return scores[-1] or []


def personal_best(scores):
    return sorted(scores)[-1] or []


def personal_top_three(scores):
    top3 = []
    scores_copy = scores.copy()
    while True:
        if len(top3) == 3 or len(scores_copy) == 0:
            break
        top_score = personal_best(scores_copy)
        top_score_index = scores_copy.index(top_score)
        top3.append(scores_copy.pop(top_score_index))
    return top3
