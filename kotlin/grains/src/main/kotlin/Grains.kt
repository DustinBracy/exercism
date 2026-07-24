import java.math.BigInteger

object Board {

    fun getGrainCountForSquare(number: Int): BigInteger {
        require(number in 1..64) {"Input must be between 1 and 64"}

        var count = 1.toBigInteger()
        for (i in 2..number) {
            count *= 2.toBigInteger()
        }
        return count
    }

    fun getTotalGrainCount(): BigInteger {
        var total = 0.toBigInteger()
        for (i in 1..64) {
            total += getGrainCountForSquare(i)
        }
        return total
    }
}
