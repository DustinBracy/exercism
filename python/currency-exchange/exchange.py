"""Submission for Exercism - Python Track - Currency Exchange"""

import math


def exchange_money(budget: float, exchange_rate: float) -> float:
    """

    :param budget: float - amount of money you are planning to exchange.
    :param exchange_rate: float - unit value of the foreign currency.
    :return: float - exchanged value of the foreign currency you can receive.
    """

    return budget / exchange_rate


def get_change(budget: float, exchanging_value: int) -> float:
    """

    :param budget: float - amount of money you own.
    :param exchanging_value: int - amount of your money you want to exchange now.
    :return: float - amount left of your starting currency after exchanging.
    """

    return budget - exchanging_value


def get_value_of_bills(denomination: int, number_of_bills: int) -> int:
    """

    :param denomination: int - the value of a bill.
    :param number_of_bills: int - amount of bills you received.
    :return: int - total value of bills you now have.
    """

    return denomination * number_of_bills


def get_number_of_bills(budget: float, denomination: int) -> int:
    """

    :param budget: float - the amount of money you are planning to exchange.
    :param denomination: int - the value of a single bill.
    :return: int - number of bills after exchanging all your money.
    """

    return math.floor(budget / denomination)


def get_real_exchange_rate(exchange_rate: float, spread: int):
    """
    Returns actual exchange rate after accounting for spread.
    :param exchange_rate: float - unit value of the foreign currency.
    :param spread: int - percentage that is taken as an exchange fee.
    :return: float - unit value of the foreign currency after spread.
    """
    return exchange_rate * (1 + spread / 100)


def exchangeable_value(
    budget: float, exchange_rate: float, spread: int, denomination: int
) -> int:
    """

    :param budget: float - the amount of your money you are planning to exchange.
    :param exchange_rate: float - the unit value of the foreign currency.
    :param spread: int - percentage that is taken as an exchange fee.
    :param denomination: int - the value of a single bill.
    :return: int - maximum value you can get.
    """
    real_exchange_rate = get_real_exchange_rate(exchange_rate, spread)
    new_money = exchange_money(budget, real_exchange_rate)
    bills = get_number_of_bills(new_money, denomination)
    return bills * denomination


def non_exchangeable_value(
    budget: float, exchange_rate: float, spread: int, denomination: int
) -> int:
    """

    :param budget: float - the amount of your money you are planning to exchange.
    :param exchange_rate: float - the unit value of the foreign currency.
    :param spread: int - percentage that is taken as an exchange fee.
    :param denomination: int - the value of a single bill.
    :return: int non-exchangeable value.
    """
    real_exchange_rate = get_real_exchange_rate(exchange_rate, spread)
    new_money = exchange_money(budget, real_exchange_rate)
    exchangeable_money = exchangeable_value(budget, exchange_rate, spread, denomination)
    return math.floor(new_money - exchangeable_money)
