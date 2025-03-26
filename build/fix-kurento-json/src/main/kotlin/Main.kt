import com.google.gson.Gson
import com.google.gson.GsonBuilder
import com.google.gson.JsonObject
import java.io.FileReader
import java.io.FileWriter
import java.io.PrintWriter
import java.nio.file.Files
import java.nio.file.Path
import java.nio.file.Paths

fun main(args: Array<String>) {
    val cwd = Paths.get("").toAbsolutePath()
    println("Working Directory = $cwd")

    val input1 = cwd.resolve("../kms-elements/src/server/interface").normalize()
    println("Search Directory = $input1")

    val output1 = cwd.resolve("../kms-elements-valid-json").normalize()

    fixFiles(input1, output1)

    val input2 = cwd.resolve("../kms-core/src/server/interface/core.kmd.json").normalize()
    println("Search Directory = $input2")

    val output2 = cwd.resolve("../kms-core-valid-json").normalize()

    fixFiles(input2, output2)
}

fun fixFiles(input: Path, output: Path) {
    val gsonReader = Gson();
    val gsonWriter = GsonBuilder()
        .setPrettyPrinting()
        .create()

    Files.walk(input).use { paths ->
        paths.filter { Files.isRegularFile(it) && it.fileName.toString().endsWith(".kmd.json") }
            .forEach {
                val o = gsonReader.fromJson(FileReader(it.toFile()), JsonObject::class.java)
                val newFileName = output.resolve(it.fileName)
                PrintWriter(FileWriter(newFileName.toFile()))
                    .use { newFile -> newFile.write(gsonWriter.toJson(o)) }
            }
    }
}