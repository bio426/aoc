const std = @import("std");

const y15 = @import("y15.zig");
const y16 = @import("y16.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const file = try std.fs.cwd().openFile("inputs/05.txt", .{});
    defer file.close();

    const readBuff = try file.readToEndAlloc(allocator, 1024 * 1024);
    defer allocator.free(readBuff);

    const res = y15.solutionD5P2(readBuff);
    std.debug.print("-> {!d}\n", .{res});

    const argsRaw = try std.process.argsAlloc(allocator);
    defer std.process.argsFree(allocator, argsRaw);

    const args = argsRaw[1..argsRaw.len];
    const myRes = processArgs(args) catch |err| {
        switch (err) {
            CliError.NoOptionProvided => {
                std.debug.print("arguments missing\n", .{});
            },
            else => {},
        }
    };
    std.debug.print("-> {!d}\n", .{myRes});
}

const CliError = error{ NoOptionProvided, InvalidInteger };

fn processArgs(args: [][:0]u8) ![3]i32 {
    var argYear: i32 = 0;
    var argDay: i32 = 0;
    var argPart: i32 = 0;

    for (args, 0..) |arg, idx| {
        // dont parse the last arg
        if (idx == args.len - 1) break;

        if (std.mem.eql(u8, arg, "--year")) {
            const val = std.fmt.parseInt(i32, args[idx + 1], 10) catch return CliError.InvalidInteger;
            argYear = val;
        }

        if (std.mem.eql(u8, arg, "--day")) {
            const val = std.fmt.parseInt(i32, args[idx + 1], 10) catch return CliError.InvalidInteger;
            argDay = val;
        }

        if (std.mem.eql(u8, arg, "--part")) {
            const val = std.fmt.parseInt(i32, args[idx + 1], 10) catch return CliError.InvalidInteger;
            argPart = val;
        }
    }

    if (argYear == 0 or argDay == 0 or argPart == 0) return CliError.NoOptionProvided;

    return [_]i32{ argYear, argDay, argPart };
}

test processArgs {
    const dummy = [_][:0]u8{ "--year", "15", "--day", "1", "--part", "1" };
    const res = try processArgs(dummy);
    try std.testing.expect(res[0] == 15);
    try std.testing.expect(res[1] == 1);
    try std.testing.expect(res[2] == 1);
}

fn getResponse(args: [3]i32) !i32 {
    var res: i32 = 0;
    res += 1;

    switch (args[0]) {
        15 => {
            switch (args[1]) {
                1 => {
                    if (args[2] == 1) {
                        // res = try y15.solutionD2P1(input);
                        std.debug.print("15-1-1\n", .{});
                    } else {
                        std.debug.print("15-1-2\n", .{});
                    }
                },
                2 => {
                    if (args[2] == 1) {
                        std.debug.print("15-2-1\n", .{});
                    } else {
                        std.debug.print("15-2-2\n", .{});
                    }
                },
                else => {
                    std.debug.print("unknown day\n", .{});
                },
            }
        },
        16 => {
            switch (args[1]) {
                1 => {
                    if (args[2] == 1) {
                        std.debug.print("16-1-1\n", .{});
                    } else {
                        std.debug.print("16-1-2\n", .{});
                    }
                },
                2 => {
                    if (args[2] == 1) {
                        std.debug.print("16-2-1\n", .{});
                    } else {
                        std.debug.print("16-2-2\n", .{});
                    }
                },

                else => {
                    std.debug.print("unknown year\n", .{});
                },
            }
        },
        else => {
            std.debug.print("unknown year\n", .{});
        },
    }
    return res;
}
