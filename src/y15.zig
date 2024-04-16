const std = @import("std");

pub fn solutionD1P1(input: []u8) i32 {
    var floor: i32 = 0;
    for (input) |char| {
        if (char == '(') {
            floor += 1;
        } else {
            floor -= 1;
        }
    }

    return floor;
}

pub fn solutionD1P2(input: []u8) i32 {
    var floor: i32 = 0;
    var pos: i32 = 0;
    for (input) |char| {
        pos += 1;
        if (char == '(') {
            floor += 1;
        } else {
            floor -= 1;
        }
        if (floor == -1) break;
    }

    return pos;
}

pub fn solutionD2P1(input: []u8) i32 {
    var total: i32 = 0;
    var linesIter = std.mem.split(u8, input, "\n");

    while (linesIter.next()) |line| {
        // to handle last empty string
        if (line.len == 0) continue;

        var dimensions: [3]i32 = undefined;
        var idx: usize = 0;
        var lineIter = std.mem.split(u8, line, "x");
        while (lineIter.next()) |dimension| {
            if (std.fmt.parseInt(i32, dimension, 10)) |parsedInt| {
                dimensions[idx] = parsedInt;
            } else |err| {
                std.debug.print("err: {!}\n", .{err});
                break;
            }
            idx += 1;
        }
        idx = 0;

        const s1 = dimensions[0] * dimensions[1];
        const s2 = dimensions[1] * dimensions[2];
        const s3 = dimensions[2] * dimensions[0];

        const boxArea = 2 * s1 + 2 * s2 + 2 * s3;
        const extra = @min(@min(s1, s2), s3);

        total += boxArea + extra;
    }

    return total;
}

pub fn solutionD2P2(input: []u8) i32 {
    var total: i32 = 0;
    var linesIter = std.mem.split(u8, input, "\n");

    while (linesIter.next()) |line| {
        // to handle last empty string
        if (line.len == 0) continue;

        var dimensions: [3]i32 = undefined;
        var idx: usize = 0;
        var lineIter = std.mem.split(u8, line, "x");
        while (lineIter.next()) |dimension| {
            if (std.fmt.parseInt(i32, dimension, 10)) |parsedInt| {
                dimensions[idx] = parsedInt;
            } else |err| {
                std.debug.print("err: {!}\n", .{err});
                break;
            }
            idx += 1;
        }
        idx = 0;

        const smallestSide: i32 = @min(@min(dimensions[0], dimensions[1]), dimensions[2]);
        var smallerSide: i32 = 0;
        if (smallestSide == dimensions[0]) {
            smallerSide = @min(dimensions[1], dimensions[2]);
        } else if (smallestSide == dimensions[1]) {
            smallerSide = @min(dimensions[0], dimensions[2]);
        } else {
            smallerSide = @min(dimensions[0], dimensions[1]);
        }

        const perimeter = 2 * smallestSide + 2 * smallerSide;
        const bow = dimensions[0] * dimensions[1] * dimensions[2];

        total += perimeter + bow;
    }

    return total;
}

pub fn solutionD3P1(input: []u8) i32 {
    var floor: i32 = 0;
    var pos: i32 = 0;
    for (input) |char| {
        pos += 1;
        if (char == '(') {
            floor += 1;
        } else {
            floor -= 1;
        }
        if (floor == -1) break;
    }

    return pos;
}
