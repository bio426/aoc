const std = @import("std");

const y15 = @import("y15.zig");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const file = try std.fs.cwd().openFile("inputs/05.txt", .{});
    defer file.close();

    const readBuff = try file.readToEndAlloc(allocator, 1024 * 1024);
    defer allocator.free(readBuff);

    const res = y15.solutionD4P1(readBuff);
    std.debug.print("-> {!d}\n", .{res});
}
