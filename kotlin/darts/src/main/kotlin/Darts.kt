import kotlin.math.sqrt

object Darts {
    fun distance(x:Number, y: Number): Double {
        val dx = x.toDouble()
        val dy = y.toDouble()
        return sqrt(dx*dx + dy*dy)
    }

    fun score(x:Number, y: Number): Int {
        val r = distance(x,y)
        val points = when {
            r <= 1 -> 10
            r <= 5 -> 5
            r <= 10 -> 1
            else -> 0
        }
        return points
    }
}
