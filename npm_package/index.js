const path = require("path");
const { platform, arch } = require("process");
const { spawn } = require("child_process");

function getBinaryPath() {
    const platformMap = {
        win32: "win32",
        linux: "linux",
        darwin: "darwin",
    };
    const archMap = {
        x64: "x64",
        ia32: "ia32",
        arm64: "arm64",
    };
    const binName = platform === "win32" ? "lekalo.exe" : "lekalo";
    return path.join(
        __dirname,
        "bin",
        `${platformMap[platform]}-${archMap[arch]}`,
        binName,
    );
}

function runCLI(args = []) {
    const binaryPath = getBinaryPath();
    const child = spawn(binaryPath, args, { stdio: "inherit" });

    child.on("error", (err) => {
        console.error("Binary executed error: ", err);
        process.exit(1);
    });

    child.on("exit", (code) => {
        process.exit(code || 0);
    });
}

const args = process.argv.slice(2);
runCLI(args);
