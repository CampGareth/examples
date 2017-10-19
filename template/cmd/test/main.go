package main

import (
    "os"
    "xcl"
)

func main() {
    // Allocate a 'world' for interacting with kernels
    world := xcl.NewWorld()
    defer world.Release()

    // Import the kernel.
    // Right now these two identifiers are hard coded as an output from the build process
    krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
    defer krnl.Release()

    // Allocate a space in the shared memory to store the results from the FPGA
    buff := world.Malloc(xcl.WriteOnly, <size here>)
    defer buff.Free()

    // Create/get data and pass arguments to the kernel as required. These could be small pieces of data,
    // pointers to memory, data lengths so the Kernel knows what to expect. This all depends on your project.
    // We have passed three arguments here, you can pass more as neccessary

    // First argument
    krnl.SetArg(0, <first>)
    // Second argument
    krnl.SetArg(1, <second>)
    // Third argument
    krnl.SetMemoryArg(2, <third>)

    // Run the kernel with the supplied arguments. This is the same for all projects.
    // The arguments ``(1, 1, 1)`` relate to x, y, z co-ordinates and correspond to our current
    // underlying technology.
    krnl.Run(1, 1, 1)

    // Display/use the results returned from the FPGA as required!

    ...

}
