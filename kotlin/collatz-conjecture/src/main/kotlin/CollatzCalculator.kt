object CollatzCalculator {
    fun computeStepCount(start: Int): Int {
        if (start < 1) throw IllegalArgumentException("Start must be a positive number")
        var iterations = 0
        var next = start
        while (next != 1) {
            if (next % 2 == 0) {
                next /= 2
            } else {
                next = (next * 3) + 1
            }
            iterations++
        }
        return iterations
    }
}