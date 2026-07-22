object EliudsEggs {

    fun eggCount(number: Int): Int{
        var eggs = 0
        var num = number
        while (num > 0) {
            eggs += num % 2
            num /= 2
        }
        return eggs
    }
}
